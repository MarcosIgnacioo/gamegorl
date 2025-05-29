package cpu

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
	RRA
	RLA
	RRCA
	RRLA
	CPL
	BIT
	RESET
	SET
	SRL
	RR
	RL
	RRC
	RLC
	SRA
	SLA
	SWAP
)

// remove the target field please and probably should hardcode the other operations because yeah uwu
func (cpu *CPU) execute(instruction Instruction, target ArithmeticTarget) {
	switch instruction {
	case ADDHL:
		{
			panic("not implemented yet because memory doesnt exist yet")
		}
	case ADD:
		{
			cpu.add(target)
		}
	case ADC:
		{
			cpu.addWithCarry(target)
		}
	case SUB:
		{
			cpu.sub(target)
		}
	case SBC:
		{
			cpu.subWithCarry(target)
		}
	case AND:
		{
			cpu.and(target)
		}
	case OR:
		{
			cpu.or(target)
		}
	case XOR:
		{
			cpu.xor(target)
		}
	case CP:
		{
			cpu.cmp(target)
		}
	case INC:
		{
			cpu.inc(target)
		}
	case DEC:
		{
			cpu.dec(target)
		}
	case CCF:
		{
			cpu.ccf()
		}
	case SCF:
		{
			cpu.scf()
		}
	case RRA:
		{
			// here i should get the value with the funcitons instead of letting them do the calculations and setting up the value, is just too much for them to handle they will get tired
			cpu.rotateRightThroughCarry(true, A)
		}
	case RLA:
		{
			cpu.rotateLeftThroughCarry(true, A)
		}
	}
}

// this updates the a register always, might refactor the registers to be an array please please please?
// we also do something about the half carry weird here cause we use the a register always which aint bad because all the math operations are performed there but be careful also, cause some of them might not do it
// unrefactor this cause this might break things hehe
func (cpu *CPU) updateFlags(value, new_value reg, did_overflow, is_substracting bool) {
	cpu.regs.f.zero = new_value == 0
	cpu.regs.f.substract = is_substracting
	cpu.regs.f.half_carry = (cpu.regs.a&0xF)+(value&0xF) > 0xF
	cpu.regs.f.carry = did_overflow
}

type ArithmeticTarget uint8

const (
	A ArithmeticTarget = iota
	B
	C
	D
	E
	H
	L
)

func (cpu *CPU) getRegisterValueByTarget(target ArithmeticTarget) reg {
	var target_reg reg
	// cpu.regs[target] in c lovely c i would have done this :3
	switch target {
	case A:
		{
			target_reg = cpu.regs.a
		}
	case B:
		{
			target_reg = cpu.regs.b
		}
	case C:
		{
			target_reg = cpu.regs.c
		}
	case D:
		{
			target_reg = cpu.regs.d
		}
	case E:
		{
			target_reg = cpu.regs.e
		}
	case H:
		{
			target_reg = cpu.regs.h
		}
	case L:
		{
			target_reg = cpu.regs.l
		}
	}
	return target_reg
}

func (cpu *CPU) getRegisterByTarget(target ArithmeticTarget) *reg {
	var target_reg *reg
	// cpu.regs[target] in c lovely c i would have done this :3
	switch target {
	case A:
		{
			target_reg = &cpu.regs.a
		}
	case B:
		{
			target_reg = &cpu.regs.b
		}
	case C:
		{
			target_reg = &cpu.regs.c
		}
	case D:
		{
			target_reg = &cpu.regs.d
		}
	case E:
		{
			target_reg = &cpu.regs.e
		}
	case H:
		{
			target_reg = &cpu.regs.h
		}
	case L:
		{
			target_reg = &cpu.regs.l
		}
	}
	return target_reg
}

func (cpu *CPU) add(target ArithmeticTarget) {
	value := cpu.getRegisterValueByTarget(target)
	new_value, did_overflow := cpu.regs.a.overflowingAdd(value)
	cpu.updateFlags(value, new_value, did_overflow, false)
	cpu.regs.a = new_value
}

func (cpu *CPU) addWithCarry(target ArithmeticTarget) {
	value := cpu.getRegisterValueByTarget(target)
	new_value, did_overflow := cpu.regs.a.overflowingAdd(value)
	if did_overflow {
		new_value += 1
	}
	cpu.updateFlags(value, new_value, did_overflow, false)
	cpu.regs.a = new_value
}

func (cpu *CPU) sub(target ArithmeticTarget) {
	value := cpu.getRegisterValueByTarget(target)
	new_value, did_overflow := cpu.regs.a.overflowingSub(value)
	cpu.updateFlags(value, new_value, did_overflow, true)
	cpu.regs.a = new_value
}

func (cpu *CPU) subWithCarry(target ArithmeticTarget) {
	value := cpu.getRegisterValueByTarget(target)
	new_value, did_overflow := cpu.regs.a.overflowingSub(value)
	if did_overflow {
		new_value -= 1
	}
	cpu.updateFlags(value, new_value, did_overflow, true)
	cpu.regs.a = new_value
}

func (cpu *CPU) and(target ArithmeticTarget) {
	value := cpu.getRegisterValueByTarget(target)
	new_value := cpu.regs.a & value
	cpu.updateFlags(value, new_value, false, false)
	cpu.regs.a = new_value
}

func (cpu *CPU) or(target ArithmeticTarget) {
	value := cpu.getRegisterValueByTarget(target)
	new_value := cpu.regs.a | value
	cpu.updateFlags(value, new_value, false, false)
	cpu.regs.a = new_value
}

func (cpu *CPU) xor(target ArithmeticTarget) {
	value := cpu.getRegisterValueByTarget(target)
	new_value := cpu.regs.a ^ value
	cpu.updateFlags(value, new_value, false, false)
	cpu.regs.a = new_value
}

func (cpu *CPU) cmp(target ArithmeticTarget) {
	value := cpu.getRegisterValueByTarget(target)
	new_value, did_overflow := cpu.regs.a.overflowingSub(value)
	cpu.updateFlags(value, new_value, did_overflow, true)
}

func (cpu *CPU) inc(target ArithmeticTarget) {
	register := cpu.getRegisterByTarget(target)
	value := reg(1)
	new_value, did_overflow := register.overflowingAdd(value)
	cpu.updateFlags(value, new_value, did_overflow, false)
	*register = new_value
}

func (cpu *CPU) dec(target ArithmeticTarget) {
	register := cpu.getRegisterByTarget(target)
	value := reg(1)
	new_value, did_overflow := register.overflowingSub(value)
	cpu.updateFlags(value, new_value, did_overflow, true)
	*register = new_value
}

// flips the carry flag and clears the *substract* and *half_carry* *flags*
func (cpu *CPU) ccf() {
	cpu.regs.f.half_carry = false
	cpu.regs.f.substract = false
	cpu.regs.f.carry = !cpu.regs.f.carry
}

// sets the *carry* flag to true and clears the *substract* and *half_carry* *flags*
func (cpu *CPU) scf() {
	cpu.regs.f.half_carry = false
	cpu.regs.f.substract = false
	cpu.regs.f.carry = true
}

// shifts right the A register once
// sets the 7nth bit of the A register to
// the value in the carry flag
// sets the carry flag to the value of the 0nth bit of
// the original value of the register A
func (cpu *CPU) rotateRightThroughCarry(set_zero bool, target ArithmeticTarget) {
	var carry_bit reg
	register := cpu.getRegisterByTarget(target)
	og_value := *register

	if cpu.regs.f.carry {
		carry_bit = 1 << 7
	}
	// shift right each bit once
	*register >>= 1
	// add the carry bit
	*register |= carry_bit

	cpu.regs.f.zero = set_zero && *register == 0
	cpu.regs.f.substract = false
	cpu.regs.f.half_carry = false
	// why if the lowest bit is on it
	// indicates the carry flag it
	// it doesnt make any sense
	// because we save the old value of the 0 bit from the A register into
	// the carry flag
	cpu.regs.f.carry = og_value&0b1 == 0b1
}

func (cpu *CPU) rotateLeftThroughCarry(set_zero bool, target ArithmeticTarget) {
	var carry_bit reg
	register := cpu.getRegisterByTarget(target)
	og_value := *register

	if cpu.regs.f.carry {
		carry_bit = 1
	}

	*register <<= 1
	*register |= carry_bit

	cpu.regs.f.zero = set_zero && *register == 0
	cpu.regs.f.substract = false
	cpu.regs.f.half_carry = false
	// 0x80 == 10000000
	// we check if the uppest bit is on in that case we turn on the
	// carry flag
	cpu.regs.f.carry = og_value&0x80 == 0x80
}
