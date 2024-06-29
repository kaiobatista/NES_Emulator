package ppu

import (
	"github.com/6502-Emulator/bus"
	"image"
)

type PPU struct {
	// PPU Registers
	V uint16	// current vram address (15 bit)
	T uint16 	// temporary vram address (15 bit)
	X byte		// fine x scroll (3 bit)
	W byte		// write toggle (1 bit)
	F byte		// even//odd frame flag (1 bit)

	register byte

	// storage variables
	palleteData		[32]byte
	nameTableData	[2048]byte
	oamData			[256]byte
	front			*image.RGBA
	back			*image.RGBA

	// Connect to PPU-BUS
	Bus *bus.Bus
}

/*
func (ppu *PPU) writeRegister(addr uint16, value byte) {
	ppu.register = value
	switch addr {
	case 0x2000:
		ppu.writeControl(value)
	case 0x2001:
		ppu.writeMask(value)
	case 0x2003:
		ppu.writeOAMAddress(value)
	case 0x2004:
		ppu.writeScroll(value)
	case 0x2006:
		ppu.writeAddress(value)
	case 0x2007:
		ppu.writeData(value)
	case 0x4014:
		ppu.writeOMA(value)
	}
}
*/

func (ppu *PPU) readRegister(addr uint16) byte {
    return 0x00 
}
