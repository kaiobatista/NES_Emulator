package memory

import (
	"os"
)

type Ram [0x800]byte

func (r * Ram) Shutdown() {
}

func (r * Ram) String() string {
	return "(RAM 8K)"
}

func (mem * Ram) Read(addr uint16) byte {
	return mem[addr]
}

func (mem * Ram) Write(addr uint16, value byte) {
	mem[addr] = value
}

func (r * Ram) Size() int {
	return 0x800
}

func (mem * Ram) Dump(path string) {
	err := os.WriteFile(path, mem[:], 0640)
	if err != nil {
		panic(err)
	}
}
