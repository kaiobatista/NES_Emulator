package bus

import (
	"fmt"
	"github.com/6502-Emulator/memory"
)

type BusEntry struct {
	mem memory.Memory
	name string
	start uint16
	end uint16 
}

type Bus struct {
	entries []BusEntry
}

func (b *Bus) backendFor(addr uint16) (memory.Memory, error) {
	for _, be := range b.entries {
		if addr >= be.start && addr <= be.end {
			return be.mem, nil
		}
	}
	return nil, fmt.Errorf("No backend for address 0x%04X", addr)
}

func (b *Bus) Read(addr uint16) byte {
	mem, err := b.backendFor(addr)
	if err != nil {
		panic(err)
	}
	value := mem.Read(addr)
	return value
}

func (b *Bus) Read16(addr uint16) uint16 {
	lo := uint16(b.Read(addr))
	hi := uint16(b.Read(addr + 1))
	return hi << 8 | lo
}


func (b *Bus) Write(addr uint16, value byte) {
	mem, err := b.backendFor(addr)
	if err != nil {
		panic(err)
	}
	mem.Write(addr, value)
}

func (b *Bus) Write16(addr uint16, value uint16) {
	b.Write(addr, byte(value))
	b.Write(addr + 1, byte(value >> 8))
}