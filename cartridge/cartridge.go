package cartridge

type Cartridge struct {
	PGR []byte	// PGR-ROM Banks
	CHR []byte  // CHR-ROM Banks
	SRAM []byte // Save RAM
	Mapper byte
	Mirror byte
	Battery byte
}

func NewCartridge(prg, chr []byte, mapper, mirror, battery byte) *Cartridge {
	sram := make([]byte, 0x2000)
	return &Cartridge{prg, chr, sram, mapper, mirror, battery}
}