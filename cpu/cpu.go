package cpu

import (
	"fmt"
	"github.com/6502-Emulator/bus"
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
	c.PC += uint16(in.Bytes)
	c.execute(in)

}

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


