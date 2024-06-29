package cartridge

import (
	"encoding/binary"
	"os"
)

type Cartridge struct {
	PRG []byte	// PRG-ROM Banks
	CHR []byte  // CHR-ROM Banks
	SRAM []byte // Save RAM
	Mapper byte
	Mirror byte
	Battery byte
}


func NewCartridge(prg, chr []byte, mapper byte, mirror, battery byte) *Cartridge {
	sram := make([]byte, 0x2000)
	return &Cartridge{prg, chr, sram, mapper, mirror, battery}
}

func LoadFromFile(path string) (*Cartridge, error)  {
    type sHeader struct {
        name [4]byte
        prg_rom_chunks byte
        chr_rom_chunks byte
        mapper1 byte
        mapper2 byte
        prg_ram_size byte
        tv_system1 byte
        tv_system2 byte
        unused [5]byte

    }

    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    var header sHeader
    err = binary.Read(file, binary.LittleEndian, &header)
    if err != nil {
        return nil, err
    }

    nMapperID := ((header.mapper2 >> 4) << 4) | (header.mapper1 >> 4) 

    mirror := ((header.mapper1 & 1) | (header.mapper1 >> 3) & 1)
    battery := (header.mapper1 >> 1) & 1

    prg := make([]byte, int(header.prg_rom_chunks) * 16384)
    chr := make([]byte, int(header.chr_rom_chunks) * 8192)

    return NewCartridge(prg, chr, nMapperID, mirror, battery), nil
}
