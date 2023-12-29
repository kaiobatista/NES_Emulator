package cpu

import (
	"fmt"

	"github.com/6502-Emulator/bus"
)

const (
	sCarry = iota
	sZero
	sInterrupt
	sDecimal
	sBreak
	_
	sOverflow
	sNegative
)

const StackBase = 0x0100

type CPU struct {
	// Program Counter
	PC uint16

	// Stack Pointer
	SP byte

	// Accumulator Register
	A byte

	// X Register
	X byte

	// Y Register
	Y byte

	// Status Register
	SR byte

	Bus *bus.Bus
}

func (c *CPU) Reset() {
	c.PC = c.Bus.Read16(0xFFFC)
	c.SR = 0b00110100
}

func (c *CPU) Step() {
	in := ReadInstruction(c.PC, c.Bus)

	fmt.Printf("addr: %X | value: %X\n", c.PC, c.Bus.Read(c.PC))

	c.PC += uint16(in.Bytes)
	c.execute(in)

}

func (c *CPU) StackHead(offset int8) uint16 {
	return uint16(StackBase) + uint16(c.SP) + uint16(offset)
}

func (c *CPU) ResolveOperand(in Instruction) uint8 {
	switch in.addressing {
	case immediate:
		return in.Op8
	default:
		return c.Bus.Read(c.memoryAddress(in))
	}
}

func (c *CPU) branch(in Instruction) {
	relative := int8(in.Op8)
	if relative > 0 {
		c.PC += uint16(relative)
	} else {
		c.PC -= uint16(-relative)
	}
}

func (c *CPU) memoryAddress(in Instruction) uint16 {
	switch in.addressing {
	case absolute:
		return in.Op16
	case absoluteX:
		return in.Op16 + uint16(c.X)
	case absoluteY:
		return in.Op16 + uint16(c.Y)
	case indirectX:
		location := uint16(in.Op8 + c.X)
		if location == 0xFF {
			fmt.Printf("Indexed indirect high-byte not on zero page.")
		}
		return c.Bus.Read16(location)
	case indirectY:
		return c.Bus.Read16(uint16(in.Op8)) + uint16(c.Y)

	case zeropage:
		return c.Bus.Read16(uint16(in.Op8))
	case zeropageX:
		return c.Bus.Read16(uint16(in.Op8 + c.X))
	case zeropageY:
		return c.Bus.Read16(uint16(in.Op8 + c.Y))

	default:
		panic("unhandled addressing")
	}
}

func (c *CPU) getStatusInt(bit uint8) uint8 {
	return (c.SR >> bit) & 1
}

func (c *CPU) getStatus(bit uint8) bool {
	return c.getStatusInt(bit) == 1
}

func (c *CPU) setStatus(bit uint8, state bool) {
	if state {
		c.SR |= 1 << bit
	} else {
		c.SR &^= 1 << bit
	}
}

func (c *CPU) updateStatus(value uint8) {
	c.setStatus(sZero, value == 0)
	c.setStatus(sNegative, (value>>7) == 1)
}

// Most Important Functions

func (c *CPU) execute(in Instruction) {
	switch in.id {
	case adc:
		c.ADC(in)
	case and:
		c.AND(in)
	case asl:
		c.ASL(in)
	case bcc:
		c.BCC(in)
	case bcs:
		c.BCS(in)
	case beq:
		c.BEQ(in)
	case bit:
		c.BIT(in)
	case bmi:
		c.BMI(in)
	case bne:
		c.BNE(in)
	case bpl:
		c.BPL(in)
	case brk:
		c.BRK(in)
	case clc:
		c.CLC(in)
	case cld:
		c.CLD(in)
	case cli:
		c.CLI(in)
	case cmp:
		c.CMP(in)
	case cpx:
		c.CPX(in)
	case cpy:
		c.CPY(in)
	case dec:
		c.DEC(in)
	case dex:
		c.DEX(in)
	case dey:
		c.DEY(in)
	case eor:
		c.EOR(in)
	case inc:
		c.INC(in)
	case inx:
		c.INX(in)
	case iny:
		c.INY(in)
	case jmp:
		c.JMP(in)
	case jsr:
		c.JSR(in)
	case lda:
		c.LDA(in)
	case ldx:
		c.LDX(in)
	case ldy:
		c.LDY(in)
	case lsr:
		c.LSR(in)
	case nop:
		c.NOP(in)
	case ora:
		c.ORA(in)
	case pha:
		c.PHA(in)
	case pla:
		c.PLA(in)
	case rol:
		c.ROL(in)
	case ror:
		c.ROR(in)
	case rts:
		c.RTS(in)
	case sbc:
		c.SBC(in)
	case sec:
		c.SEC(in)
	case sei:
		c.SEI(in)
	case sta:
		c.STA(in)
	case stx:
		c.STX(in)
	case sty:
		c.STY(in)
	case tax:
		c.TAX(in)
	case tay:
		c.TAY(in)
	case tsx:
		c.TSX(in)
	case txa:
		c.TXA(in)
	case txs:
		c.TXS(in)
	case tya:
		c.TYA(in)
	case _end:
		c._END(in)
	default:
		panic(fmt.Sprintf("unhandled instruction: %v", in))
	}
}

func (c *CPU) ADC(in Instruction) {
	value16 := uint16(c.A) + uint16(c.ResolveOperand(in)) + uint16(c.getStatusInt(sCarry))
	c.setStatus(sCarry, value16 > 0xFF)
	c.A = uint8(value16)
	c.updateStatus(c.A)
}

func (c *CPU) AND(in Instruction) {
	c.A &= c.ResolveOperand(in)
	c.updateStatus(c.A)
}

func (c *CPU) ASL(in Instruction) {
	switch in.addressing {
	case accumulator:
		c.setStatus(sCarry, (c.A>>7) == 1)
		c.A <<= 1
		c.updateStatus(c.A)
	default:
		address := c.memoryAddress(in)
		value := c.Bus.Read(address)
		c.setStatus(sCarry, (value>>7) == 1)
		value <<= 1
		c.Bus.Write(address, value)
		c.updateStatus(value)
	}
}

func (c *CPU) BCC(in Instruction) {
	if !c.getStatus(sCarry) {
		c.branch(in)
	}
}

func (c *CPU) BCS(in Instruction) {
	if c.getStatus(sCarry) {
		c.branch(in)
	}
}

func (c *CPU) BEQ(in Instruction) {
	if c.getStatus(sZero) {
		c.branch(in)
	}
}

func (c *CPU) BIT(in Instruction) {
	value := c.ResolveOperand(in)
	c.setStatus(sZero, value & c.A == 0)
	c.setStatus(sOverflow, value & (1 << 6) != 0)
	c.setStatus(sNegative, value & (1 << 7) != 0)
}

func (c *CPU) BMI(in Instruction) {
	if c.getStatus(sNegative) {
		c.branch(in)
	}
}

func (c *CPU) BNE(in Instruction) {
	if !c.getStatus(sZero) {
		c.branch(in)
	}
}

func (c *CPU) BPL(in Instruction) {
	if !c.getStatus(sNegative) {
		c.branch(in)
	}
}

func (c *CPU) BRK(in Instruction) {
	fmt.Println("BRK: ", c)
}

func (c *CPU) CLC(in Instruction) {
	c.setStatus(sCarry, false)
}

func (c *CPU) CLD(in Instruction) {
	c.setStatus(sDecimal, false)
}

func (c *CPU) CLI(in Instruction) {
	c.setStatus(sInterrupt, false)
}

func (c *CPU) CMP(in Instruction) {
	value := c.ResolveOperand(in)
	c.setStatus(sCarry, c.A >= value)
	c.updateStatus(c.A - value)
}

func (c *CPU) CPX(in Instruction) {
	value := c.ResolveOperand(in)
	c.setStatus(sCarry, c.X >= value)
	c.updateStatus(c.X - value)
}

func (c *CPU) CPY(in Instruction) {
	value := c.ResolveOperand(in)
	c.setStatus(sCarry, c.Y >= value)
	c.updateStatus(c.X - value)
}

func (c *CPU) DEC(in Instruction) {
	addr := c.memoryAddress(in)
	value := c.Bus.Read(addr) - 1
	c.Bus.Write(addr, value)
	c.updateStatus(value)
}

func (c *CPU) DEX(in Instruction) {
	c.X--
	c.updateStatus(c.X)
}

func (c *CPU) DEY(in Instruction) {
	c.Y--
	c.updateStatus(c.Y)
}

func(c *CPU) EOR(in Instruction) {
	value := c.ResolveOperand(in)
	c.A ^= value
	c.updateStatus(c.A)
}

func (c *CPU) INC(in Instruction) {
	addr := c.memoryAddress(in)
	value := c.Bus.Read(addr) + 1
	c.Bus.Write(addr, value)
	c.updateStatus(value)
}

func (c *CPU) INX(in Instruction) {
	c.X++
	c.updateStatus(c.X)
}

func (c *CPU) INY(in Instruction) {
	c.Y++
	c.updateStatus(c.Y)
}

func (c *CPU) JMP(in Instruction) {
	c.PC = c.memoryAddress(in)
}

func (c *CPU) JSR(in Instruction) {
	c.Bus.Write16(c.StackHead(-1), c.PC - 1)
	c.SP -= 2
	c.PC = in.Op16
}

func (c *CPU) LDA(in Instruction) {
	c.A = c.ResolveOperand(in)
	c.updateStatus(c.A)
}

func (c *CPU) LDX(in Instruction) {
	c.X = c.ResolveOperand(in)
	c.updateStatus(c.X)
}

func (c *CPU) LDY(in Instruction) {
	c.Y = c.ResolveOperand(in)
	c.updateStatus(c.Y)
}

func (c *CPU) LSR(in Instruction) {
	switch in.addressing {
	case accumulator:
		c.setStatus(sCarry, c.A & 1 == 1)
		c.A >>= 1
		c.updateStatus(c.A)
	default:
		address := c.memoryAddress(in)
		value := c.Bus.Read(address)
		c.setStatus(sCarry, value&1 == 1)
		value >>= 1
		c.Bus.Write(address, value)
		c.updateStatus(value)
	}
}

func (c *CPU) NOP(in Instruction) {
}

func (c *CPU) ORA(in Instruction) {
	c.A |= c.ResolveOperand(in)
	c.updateStatus(c.A)
}

func (c *CPU) PHA(in Instruction) {
	c.Bus.Write(0x0100 + uint16(c.SP), c.A)
	c.SP--
}

func (c *CPU) PLA(in Instruction) {
	c.SP++
	c.Bus.Read(0x0100 + uint16(c.SP))
}

func (c *CPU) ROL(in Instruction) {
	carry := c.getStatusInt(sCarry)
	switch in.addressing {
	case accumulator:
		c.setStatus(sCarry, c.A >> 7 == 1)
		c.A = c.A << 1 | carry
		c.updateStatus(c.A)
	default:
		address := c.memoryAddress(in)
		value := c.Bus.Read(address)
		c.setStatus(sCarry, value>>7 == 1)
		value = value<<1 | carry
		c.Bus.Write(address, value)
		c.updateStatus(value)
	}
}

func (c *CPU) ROR(in Instruction) {
	carry := c.getStatusInt(sCarry)
	switch in.addressing {
	case accumulator:
		c.setStatus(sCarry, c.A & 1 == 1)
		c.A = c.A >> 1 | carry << 7
		c.updateStatus(c.A)
	default:
		address := c.memoryAddress(in)
		value := c.Bus.Read(address)
		c.setStatus(sCarry, value&1 == 1)
		value = value>>1 | carry<<7
		c.Bus.Write(address, value)
		c.updateStatus(value)
	}
}

func (c *CPU) RTS(in Instruction) {
	c.PC = c.Bus.Read16(c.StackHead(1))
	c.SP += 2
	c.PC++
}

func (c *CPU) SBC(in Instruction) {
	valueSigned := int16(c.A) - int16(c.ResolveOperand(in))
	if !c.getStatus(sCarry) {
		valueSigned--
	}
	c.A = uint8(valueSigned)
	// TODO: Set v flag - c.setStatus(sOverflow, ...)

	c.setStatus(sCarry, valueSigned >= 0)
	
	c.updateStatus(c.A)
}

func (c *CPU) SEC(in Instruction) {
	c.setStatus(sCarry, true)
}

func (c *CPU) SEI(in Instruction) {
	c.setStatus(sInterrupt, true)
}

func (c *CPU) STA(in Instruction) {
	c.Bus.Write(c.memoryAddress(in), c.A)
}

func (c *CPU) STX(in Instruction) {
	c.Bus.Write(c.memoryAddress(in), c.X)
}

func (c *CPU) STY(in Instruction) {
	c.Bus.Write(c.memoryAddress(in), c.Y)
}

func (c *CPU) TAX(in Instruction) {
	c.X = c.A
	c.updateStatus(c.X)
}

func (c *CPU) TAY(in Instruction) {
	c.Y = c.A
	c.updateStatus(c.Y)
}

func (c *CPU) TSX(in Instruction) {
	c.X = c.SP
	c.updateStatus(c.X)
}

func (c *CPU) TXA(in Instruction) {
	c.A = c.X
	c.updateStatus(c.A)
}

func (c *CPU) TYA(in Instruction) {
	c.A = c.Y
	c.updateStatus(c.A)
}

func (c *CPU) TXS(in Instruction) {
	c.SP = c.X
	c.updateStatus(c.SP)
}

func (c *CPU) _END(in Instruction) {
}