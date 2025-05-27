package cpu

import "fmt"

// la funcion de toBit la pude hacer porque sabia que las expressiones finales de los if statements son a lo que terminan evaluandose, por lo que probablemente haya mas cosas que no tenga la menor idea asi que tendre que buscar la referencia de lo que significa en internet o preguntandole a alguient
type Instruction uint8

const (
	ADD Instruction = iota
	ADDHL
	ADC
	SUB
	SBC
	AND
	OR
	XOR
	CP
	INC
	DEC
	CCF
	SCF
)

type CPU struct {
	regs Registers
}

// half_carry doc
// a     : 0111
// value : 1111
//
//	10110 > 1111 is half carry, else, there is not half carry

type reg uint8

type Registers struct {
	a reg
	b reg
	c reg
	d reg
	e reg
	f FlagsRegister
	h reg
	l reg
}

func (r *reg) overflowingAdd(b reg) (reg, bool) {
	var overflow bool
	result := (*r) + b
	if result < (*r) {
		overflow = true
	}
	return result, overflow
}

func (r *reg) overflowingSub(b reg) (reg, bool) {
	var overflow bool
	result := (*r) - b
	if result > (*r) {
		overflow = true
	}
	return result, overflow
}

func (r *Registers) String() string {
	return fmt.Sprintf(`
		[
			a: %d
			b: %d
			c: %d
			d: %d
			e: %d
			f: %v
			h: %d
			l: %d
		]
		`, r.a, r.b, r.c, r.d, r.e, r.f, r.h, r.l)
}

func (r *Registers) getBC() uint16 {
	return uint16(r.b)<<8 | uint16(r.c)
}

func (r *Registers) setBC(value uint16) {
	// [0101010101010101]
	// [1111111100000000]
	// [0101010100000000]
	// [0000000001010101]
	//
	// [0101010101010101]
	// [0000000011111111]

	// full := ^uint16(0)
	// b := uint8((value & full << 8) >> 8)
	// c := uint8((value & full >> 8))
	b := reg((value & 0xFF00) >> 8)
	c := reg(value & 0xFF)
	r.b = b
	r.c = c
}

type FlagsRegister struct {
	zero       bool
	substract  bool
	half_carry bool
	carry      bool
}

const (
	ZERO_FLAG_BYTE_POSITION       = 7
	SUBTRACT_FLAG_BYTE_POSITION   = 6
	HALF_CARRY_FLAG_BYTE_POSITION = 5
	CARRY_FLAG_BYTE_POSITION      = 4
)

// [11111111]
// [00010000]
// [00010000]
// [00100000]
// [01000000]
// [10000000]
const (
	ZERO_FLAG       = (1) << ZERO_FLAG_BYTE_POSITION
	SUBTRACT_FLAG   = (1) << SUBTRACT_FLAG_BYTE_POSITION
	HALF_CARRY_FLAG = (1) << HALF_CARRY_FLAG_BYTE_POSITION
	CARRY_FLAG      = (1) << CARRY_FLAG_BYTE_POSITION
)

func (f *FlagsRegister) toBit() uint8 {

	var flag uint8

	if f.zero {
		flag |= ZERO_FLAG
	}
	if f.substract {
		flag |= SUBTRACT_FLAG
	}
	if f.half_carry {
		flag |= HALF_CARRY_FLAG
	}
	if f.carry {
		flag |= CARRY_FLAG
	}

	return flag

}

func toFlagReg(value uint8) FlagsRegister {

	var flags FlagsRegister
	if value&ZERO_FLAG == ZERO_FLAG {
		flags.zero = true
	}

	if value&SUBTRACT_FLAG == SUBTRACT_FLAG {
		flags.substract = true
	}

	if value&HALF_CARRY_FLAG == HALF_CARRY_FLAG {
		flags.half_carry = true
	}

	if value&CARRY_FLAG == CARRY_FLAG {
		flags.carry = true
	}

	return flags

}
