package memory

import (
	"os"
)

type Rom struct {
	name string
	size int
	data []byte
}

func (r *Rom) Shutdown() {
}


func CreateRom(name string, data []byte) (*Rom, error) {
	return &Rom{name: name, size: len(data), data: data}, nil
}


func (r *Rom) Read(addr uint16) byte {
	return r.data[addr]
}

func RomFromFile(path string) (*Rom, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &Rom{name: path, size: len(data), data: data}, nil
}

func (r *Rom) Size() int {
	return r.size
}

func (r *Rom) Write(addr uint16, value byte) {
    r.data[addr] = value
}

