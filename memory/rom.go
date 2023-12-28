package memory

import (
	"io/ioutil"
)

type Rom struct {
	name string
	size int
	data []byte
}

func (r *Rom) Shutdown() {
}

func (r *Rom) Read(addr uint16) byte {
	return r.data[addr]
}

func RomFromFile(path string) (*Rom, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &Rom{name: path, size: len(data), data: data}, nil
}

func (r *Rom) Size() int {
	return r.size
}

func (r *Rom) Write(_ uint16, _ byte) {	
}

