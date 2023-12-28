package memory

import (
	"io/ioutil"
)

type Ram [0x8000]byte

func (r * Ram) Shutdown() {
}

func (r * Ram) String() string {
	return "(RAM 32K)"
}

func (mem * Ram) Read(addr uint16) byte {
	return mem[addr]
}

func (mem * Ram) Write(addr uint16, value byte) {
	mem[addr] = value
}

func (r * Ram) Size() int {
	return 0x8000
}

func (mem * Ram) Dump(path string) {
	err := ioutil.WriteFile(path, mem[:], 0640)
	if err == nil {
		panic(err)
	}
}
