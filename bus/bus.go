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

func CreateBus() (*Bus, error) {
	return &Bus{entries: make([]BusEntry, 0)}, nil
}

func (b *Bus) Attach(mem memory.Memory, name string, offset uint16) error {
	om := OffsetMemory{Offset: offset, Memory: mem}
	end := offset + uint16(mem.Size() - 1)
	entry := BusEntry{mem: om, name: name, start: offset, end: end}
	b.entries = append(b.entries, entry)
	return nil
}

func (b *Bus) backendFor(addr uint16) (memory.Memory, error) {
	for _, be := range b.entries {
		if addr >= be.start && addr <= be.end {
			return be.mem, nil
		}
	}
	return nil, fmt.Errorf("no backend for address 0x%04X", addr)
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