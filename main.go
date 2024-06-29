package main

import (
    "fmt"
	"github.com/6502-Emulator/bus"
	"github.com/6502-Emulator/cpu"
    "github.com/6502-Emulator/ppu"
	"github.com/6502-Emulator/memory"
    "github.com/6502-Emulator/cartridge"
)

func main() {
	data := make([]byte, 0x7FFF)

	rom, _ := memory.CreateRom("rom", data)
	ram := &memory.Ram{}

	addrBus, _ := bus.CreateBus()
	addrBus.Attach(ram, "ram", 0x0000)
	addrBus.Attach(rom, "rom", 0x8000)
    addrBus.Write16(0xFFFC, 0x8000)

	cpu := &cpu.CPU{Bus: addrBus}
    ppu := &ppu.PPU{Bus: addrBus}
  
    fmt.Printf("0x%04X\n", ppu.T)

    cpu.Reset()
    
	for i := 0; i < 41; i++ {
		cpu.Step()
	}

    cart, err := cartridge.LoadFromFile("./donkey_kong.nes")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(len(cart.CHR))
}
