package bus

import (
	"fmt"
	"github.com/6502-Emulator/memory"
)

type OffsetMemory struct {
	Offset uint16
	memory.Memory
}

func (om OffsetMemory) Read(addr uint16) byte {
	return om.Memory.Read(addr - om.Offset)
}


func (om OffsetMemory) String() string {
	return fmt.Sprintf("OffsetMemory(%v)", om.Memory)
}

func (om OffsetMemory) Write(addr uint16, value byte) {
	om.Memory.Write(addr - om.Offset, value)
}