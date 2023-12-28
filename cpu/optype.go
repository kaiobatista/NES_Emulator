package cpu

const (
	_ = iota
	absolute 
	absoluteX
	absoluteY
	accumulator
	immediate
	implied
	indirect
	indirectX
	indirectY
	relative
	zeropage
	zeropageX
	zeropageY
)

var addressingNames = [...]string{
	"",
	"absolute",
	"absoluteX",
	"absoluteY",
	"accumulator",
	"immediate",
	"implied",
	"(indirect)",
	"(indirect, X)",
	"(indirect, Y)",
	"relative",
	"zeropage",
	"zeropageX",
	"zeropageY",
}

const (
	_ = iota
	adc
	and
	asl
	bcc
	bcs
	beq
	bit
	bmi
	bne
	bpl
	brk
	bvc
	bvs
	clc
	cld
	cli
	clv
	cmp
	cpx
	cpy
	dec
	dex
	dey
	eor
	inc
	inx
	iny
	jmp
	jsr
	lda
	ldx
	ldy
	lsr
	nop
	ora
	pha
	php
	pla
	plp
	rol
	ror
	rti
	rts
	sbc
	sec
	sed
	sei
	sta
	stx
	sty
	tax
	tay
	tsx
	txa
	txs
	tya
	_end
)

var instructionNames = [...]string{
	"",
	"ADC",
	"AND",
	"ASL",
	"BCC",
	"BCS",
	"BEQ",
	"BIT",
	"BMI",
	"BNE",
	"BPL",
	"BRK",
	"BVC",
	"BVS",
	"CLC",
	"CLD",
	"CLI",
	"CLV",
	"CMP",
	"CPX",
	"CPY",
	"DEC",
	"DEX",
	"DEY",
	"EOR",
	"INC",
	"INX",
	"INY",
	"JMP",
	"JSR",
	"LDA",
	"LDX",
	"LDY",
	"LSR",
	"NOP",
	"ORA",
	"PHA",
	"PHP",
	"PLA",
	"PLP",
	"ROL",
	"ROR",
	"RTI",
	"RTS",
	"SBC",
	"SEC",
	"SED",
	"SEI",
	"STA",
	"STX",
	"STY",
	"TAX",
	"TAY",
	"TSX",
	"TXA",
	"TXS",
	"TYA",
	"_END",
}


type OpType struct {
	Opcode byte
	id uint8
	addressing uint8
	Bytes uint8
	Cycles uint8
}

func (ot OpType) IsAbsolute() bool {
	return ot.addressing == absolute
}


var optypes = map[uint8]OpType{
	0x69: {0x69, adc, immediate, 2, 2},
	0x65: {0x65, adc, zeropage, 2, 3},
	0x75: {0x75, adc, zeropageX, 2, 4},
	0x6D: {0x6D, adc, absolute, 3, 4},
	0x7D: {0x7D, adc, absoluteX, 3, 4},
	0x79: {0x79, adc, absoluteY, 3, 4},
	0x61: {0x61, adc, indirectX, 2, 6},
	0x71: {0x71, adc, indirectY, 2, 5},
	0x29: {0x29, and, immediate, 2, 2},
	0x25: {0x25, and, zeropage, 2, 3},
	0x35: {0x35, and, zeropageX, 2, 4},
	0x2D: {0x2D, and, absolute, 3, 4},
	0x3D: {0x3D, and, absoluteX, 3, 4},
	0x39: {0x39, and, absoluteY, 3, 4},
	0x21: {0x21, and, indirectX, 2, 6},
	0x31: {0x31, and, indirectY, 2, 5},
	0x0A: {0x0A, asl, accumulator, 1, 2},
	0x06: {0x06, asl, zeropage, 2, 5},
	0x16: {0x16, asl, zeropageX, 2, 6},
	0x0E: {0x0E, asl, absolute, 3, 6},
	0x1E: {0x1E, asl, absoluteX, 3, 7},
	0x90: {0x90, bcc, relative, 2, 2},
	0xB0: {0xB0, bcs, relative, 2, 2},
	0xF0: {0xF0, beq, relative, 2, 2},
	0x24: {0x24, bit, zeropage, 2, 3},
	0x2C: {0x2C, bit, absolute, 3, 4},
	0x30: {0x30, bmi, relative, 2, 2},
	0xD0: {0xD0, bne, relative, 2, 2},
	0x10: {0x10, bpl, relative, 2, 2},
	0x00: {0x00, brk, implied, 1, 7},
	0x50: {0x50, bvc, relative, 2, 2},
	0x70: {0x70, bvs, relative, 2, 2},
	0x18: {0x18, clc, implied, 1, 2},
	0xD8: {0xD8, cld, implied, 1, 2},
	0x58: {0x58, cli, implied, 1, 2},
	0xB8: {0xB8, clv, implied, 1, 2},
	0xC9: {0xC9, cmp, immediate, 2, 2},
	0xC5: {0xC5, cmp, zeropage, 2, 3},
	0xD5: {0xD5, cmp, zeropageX, 2, 4},
	0xCD: {0xCD, cmp, absolute, 3, 4},
	0xDD: {0xDD, cmp, absoluteX, 3, 4},
	0xD9: {0xD9, cmp, absoluteY, 3, 4},
	0xC1: {0xC1, cmp, indirectX, 2, 6},
	0xD1: {0xD1, cmp, indirectY, 2, 5},
	0xE0: {0xE0, cpx, immediate, 2, 2},
	0xE4: {0xE4, cpx, zeropage, 2, 3},
	0xEC: {0xEC, cpx, absolute, 3, 4},
	0xC0: {0xC0, cpy, immediate, 2, 2},
	0xC4: {0xC4, cpy, zeropage, 2, 3},
	0xCC: {0xCC, cpy, absolute, 3, 4},
	0xC6: {0xC6, dec, zeropage, 2, 5},
	0xD6: {0xD6, dec, zeropageX, 2, 6},
	0xCE: {0xCE, dec, absolute, 3, 3},
	0xDE: {0xDE, dec, absoluteX, 3, 7},
	0xCA: {0xCA, dex, implied, 1, 2},
	0x88: {0x88, dey, implied, 1, 2},
	0x49: {0x49, eor, immediate, 2, 2},
	0x45: {0x45, eor, zeropage, 2, 3},
	0x55: {0x55, eor, zeropageX, 2, 4},
	0x4D: {0x4D, eor, absolute, 3, 4},
	0x5D: {0x5D, eor, absoluteX, 3, 4},
	0x59: {0x59, eor, absoluteY, 3, 4},
	0x41: {0x41, eor, indirectX, 2, 6},
	0x51: {0x51, eor, indirectY, 2, 5},
	0xE6: {0xE6, inc, zeropage, 2, 5},
	0xF6: {0xF6, inc, zeropageX, 2, 6},
	0xEE: {0xEE, inc, absolute, 3, 6},
	0xFE: {0xFE, inc, absoluteX, 3, 7},
	0xE8: {0xE8, inx, implied, 1, 2},
	0xC8: {0xC8, iny, implied, 1, 2},
	0x4C: {0x4C, jmp, absolute, 3, 3},
	0x6C: {0x6C, jmp, indirect, 3, 5},
	0x20: {0x20, jsr, absolute, 3, 6},
	0xA9: {0xA9, lda, immediate, 2, 2},
	0xA5: {0xA5, lda, zeropage, 2, 3},
	0xB5: {0xB5, lda, zeropageX, 2, 4},
	0xAD: {0xAD, lda, absolute, 3, 4},
	0xBD: {0xBD, lda, absoluteX, 3, 4},
	0xB9: {0xB9, lda, absoluteY, 3, 4},
	0xA1: {0xA1, lda, indirectX, 2, 6},
	0xB1: {0xB1, lda, indirectY, 2, 5},
	0xA2: {0xA2, ldx, immediate, 2, 2},
	0xA6: {0xA6, ldx, zeropage, 2, 3},
	0xB6: {0xB6, ldx, zeropageY, 2, 4},
	0xAE: {0xAE, ldx, absolute, 3, 4},
	0xBE: {0xBE, ldx, absoluteY, 3, 4},
	0xA0: {0xA0, ldy, immediate, 2, 2},
	0xA4: {0xA4, ldy, zeropage, 2, 3},
	0xB4: {0xB4, ldy, zeropageX, 2, 4},
	0xAC: {0xAC, ldy, absolute, 3, 4},
	0xBC: {0xBC, ldy, absoluteX, 3, 4},
	0x4A: {0x4A, lsr, accumulator, 1, 2},
	0x46: {0x46, lsr, zeropage, 2, 5},
	0x56: {0x56, lsr, zeropageX, 2, 6},
	0x4E: {0x4E, lsr, absolute, 3, 6},
	0x5E: {0x5E, lsr, absoluteX, 3, 7},
	0xEA: {0xEA, nop, implied, 1, 2},
	0x09: {0x09, ora, immediate, 2, 2},
	0x05: {0x05, ora, zeropage, 2, 3},
	0x15: {0x15, ora, zeropageX, 2, 4},
	0x0D: {0x0D, ora, absolute, 3, 4},
	0x1D: {0x1D, ora, absoluteX, 3, 4},
	0x19: {0x19, ora, absoluteY, 3, 4},
	0x01: {0x01, ora, indirectX, 2, 6},
	0x11: {0x11, ora, indirectY, 2, 5},
	0x48: {0x48, pha, implied, 1, 3},
	0x08: {0x08, php, implied, 1, 3},
	0x68: {0x68, pla, implied, 1, 4},
	0x28: {0x28, php, implied, 1, 4},
	0x2A: {0x2A, rol, accumulator, 1, 2},
	0x26: {0x26, rol, zeropage, 2, 5},
	0x36: {0x36, rol, zeropageX, 2, 6},
	0x2E: {0x2E, rol, absolute, 3, 6},
	0x3E: {0x3E, rol, absoluteX, 3, 7},
	0x6A: {0x6A, ror, accumulator, 1, 2},
	0x66: {0x66, ror, zeropage, 2, 5},
	0x76: {0x76, ror, zeropageX, 2, 6},
	0x6E: {0x6E, ror, absolute, 3, 6},
	0x7E: {0x7E, ror, absoluteX, 3, 7},
	0x40: {0x40, rti, implied, 1, 6},
	0x60: {0x60, rts, implied, 1, 6},
	0xE9: {0xE9, sbc, immediate, 2, 2},
	0xE5: {0xE5, sbc, zeropage, 2, 3},
	0xF5: {0xF5, sbc, zeropageX, 2, 4},
	0xED: {0xED, sbc, absolute, 3, 4},
	0xFD: {0xFD, sbc, absoluteX, 3, 4},
	0xF9: {0xF9, sbc, absoluteY, 3, 4},
	0xE1: {0xE1, sbc, indirectX, 2, 6},
	0xF1: {0xF1, sbc, indirectY, 2, 5},
	0x38: {0x38, sec, implied, 1, 2},
	0xF8: {0xF8, sed, implied, 1, 2},
	0x78: {0x78, sei, implied, 1, 2},
	0x85: {0x85, sta, zeropage, 2, 3},
	0x95: {0x95, sta, zeropageX, 2, 4},
	0x8D: {0x8D, sta, absolute, 3, 4},
	0x9D: {0x9D, sta, absoluteX, 3, 5},
	0x99: {0x99, sta, absoluteY, 3, 5},
	0x81: {0x81, sta, indirectX, 2, 6},
	0x91: {0x91, sta, indirectY, 2, 6},
	0x86: {0x86, stx, zeropage, 2, 3},
	0x96: {0x96, stx, zeropageY, 2, 4},
	0x8E: {0x8E, stx, absolute, 3, 4},
	0x84: {0x84, sty, zeropage, 2, 3},
	0x94: {0x94, sty, zeropageX, 2, 4},
	0x8C: {0x8C, sty, absolute, 3, 4},
	0xAA: {0xAA, tax, implied, 1, 2},
	0xA8: {0xA8, tay, implied, 1, 2},
	0xBA: {0xBA, tsx, implied, 1, 2},
	0x8A: {0x8A, txa, implied, 1, 2},
	0x9A: {0x9A, txs, implied, 1, 2},
	0x98: {0x98, tya, implied, 1, 2},
	0xFF: {0xFF, _end, implied, 1, 1},
}