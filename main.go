package main

import (
	"fmt"

	"github.com/6502-Emulator/bus"
	"github.com/6502-Emulator/cpu"
	"github.com/6502-Emulator/memory"
)

func main() {
	data := make([]byte, 0xFFF)
	data[0xFFC] = 0x00
	data[0xFFD] = 0x80 

	charRom, _ := memory.RomFromFile("rom.bin")
	kernelRom, _ := memory.CreateRom("kernal", data)

	ram := &memory.Ram{}

	addrBus, _ := bus.CreateBus()
	addrBus.Attach(ram, "ram", 0x0000)
	addrBus.Attach(charRom, "rom", 0x8000)
	addrBus.Attach(kernelRom, "kernel", 0xF000)

	cpu := &cpu.CPU{Bus: addrBus}
	
	cpu.Reset()

	for i := 0; i < 41; i++ {
		cpu.Step()
	}
	
	

	fmt.Println(cpu.Bus.Read(0x0002))

	fmt.Println("CPU: ", cpu)
}