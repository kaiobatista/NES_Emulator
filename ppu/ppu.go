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

	// storage variables
	palleteData		[32]byte
	nameTableData	[2048]byte
	oamData			[256]byte
	front			*image.RGBA
	back			*image.RGBA

	// Connect to PPU-BUS
	Bus *bus.Bus
}