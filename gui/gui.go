package main

import (
	"fmt"
	"strconv"
	"strings"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"

	"github.com/6502-Emulator/bus"
	"github.com/6502-Emulator/cpu"
	"github.com/6502-Emulator/memory"
)

var addrBus, _ = bus.CreateBus()
var cpu_ = &cpu.CPU{Bus: addrBus}

func main() {
	data := make([]byte, 0x8000)
	data[0xFFFC - 0x8000] = 0x00
	data[0xFFFD - 0x8000] = 0x80 

	code := strings.Fields("A9 05 8D 01 00 A9 08 8D 02 00 A9 00 AC 02 00 18 6D 01 00 88 D0 FA 8D 02 00")

	for i, value := range code {
		v, _ := strconv.ParseUint(value, 16, 16)
		data[i] = uint8(v)

	}
	

	ram := &memory.Ram{}
	rom, _ := memory.CreateRom("rom", data)

	addrBus.Attach(ram, "ram", 0x0000)
	addrBus.Attach(rom, "rom", 0x8000)

	cpu_.Reset()
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title: "Monitoring Emulator",
		Bounds: pixel.R(0, 0, 1600, 900),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	text01 := text.New(pixel.V(0, 880), atlas)
	text02 := text.New(pixel.V(0, 500), atlas)
	text03 := text.New(pixel.V(900, 880), atlas)
	text04 := text.New(pixel.V(900, 500), atlas)

	

	for !win.Closed() {
		win.Clear(color.RGBA{R: 30, G: 35, B: 69})

		text01.Draw(win, pixel.IM.Scaled(text01.Orig, 1.7))
		text02.Draw(win, pixel.IM.Scaled(text02.Orig, 1.7))
		text03.Draw(win, pixel.IM.Scaled(text03.Orig, 1.7))
		text04.Draw(win, pixel.IM.Scaled(text04.Orig, 2))

		if win.JustPressed(pixelgl.KeySpace) {
			cpu_.Step()
		}
		
		text01.Clear()
		text02.Clear()
		text03.Clear()
		text04.Clear()

		for i := 0x00; i <= 0x0F; i++ {
			fmt.Fprintf(text01, "$%03X0: ", i)
			fmt.Fprintf(text02, "$8%02X0: ", i)
			for j := 0x00; j <= 0x0F; j++ {
				fmt.Fprintf(text01, "%02X  ", cpu_.Bus.Read(uint16(i << 4) + uint16(j)))
				fmt.Fprintf(text02, "%02X  ", cpu_.Bus.Read(0x8000 + uint16(i << 4) + uint16(j)))
			}
			fmt.Fprintf(text01, "\n")
			fmt.Fprintf(text02, "\n")
		}
	
		fmt.Fprintf(
			text03, 
			"PC: $%04X\nA: $%02X\t[%d]\nX: $%02X\t[%d]\nY: $%02X\t[%d]", 
			cpu_.PC, cpu_.A, cpu_.A, cpu_.X, cpu_.X, cpu_.Y, cpu_.Y,
		)

		in := cpu.ReadInstruction(cpu_.PC, cpu_.Bus)

		fmt.Fprintf(
			text04, "$%04X: %s ", cpu_.PC, in.String(),
		)


		win.Update()
	}
}
