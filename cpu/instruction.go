package cpu

import (
	"fmt"
	"github.com/6502-Emulator/bus"
)

type Instruction struct {
	OpType

	Op8 uint8

	Op16 uint16
}

func (in Instruction) String() (s string) {
	switch in.Bytes {
	case 3:
		s = fmt.Sprintf("%v $%04X", in.OpType, in.Op16)
	case 2:
		s = fmt.Sprintf("%v $%02X", in.OpType, in.Op8)
	case 1:
		s = in.OpType.String()
	}
	return
}


func ReadInstruction(pc uint16, bus *bus.Bus) Instruction {
	opcode := bus.Read(pc)
	optype, ok := optypes[opcode]
	if !ok {
		panic(fmt.Sprintf("Illegal opcode $%02X at $%04X", opcode, pc))
	}
	in := Instruction{OpType: optype}
	switch in.Bytes {
	case 1: // No operand
	case 2:
		in.Op8 = bus.Read(pc + 1)
	case 3:
		in.Op16 = bus.Read16(pc + 1)
	default:
		panic(fmt.Sprintf("unhandled instruction length: %d", in.Bytes))
	}
	return in
}
