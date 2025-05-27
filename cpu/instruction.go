package cpu

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

	}
}

// this updates the a register always, might refactor the registers to be an array please please please?
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
	cpu.regs.f.zero = new_value == 0
	cpu.regs.f.substract = false
	cpu.regs.f.half_carry = (cpu.regs.a&0xF)+(value&0xF) > 0xF
	cpu.regs.f.carry = did_overflow
	cpu.regs.a = new_value
}

func (cpu *CPU) addWithCarry(target ArithmeticTarget) {
	value := cpu.getRegisterValueByTarget(target)
	new_value, did_overflow := cpu.regs.a.overflowingAdd(value)
	if did_overflow {
		new_value += 1
	}
	cpu.regs.f.zero = new_value == 0
	cpu.regs.f.substract = false
	cpu.regs.f.half_carry = (cpu.regs.a&0xF)+(value&0xF) > 0xF
	cpu.regs.f.carry = did_overflow
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
	cpu.updateFlags(value, new_value, false, true)
	cpu.regs.a = new_value
}

func (cpu *CPU) or(target ArithmeticTarget) {
	value := cpu.getRegisterValueByTarget(target)
	new_value := cpu.regs.a | value
	cpu.updateFlags(value, new_value, false, true)
	cpu.regs.a = new_value
}

func (cpu *CPU) xor(target ArithmeticTarget) {
	value := cpu.getRegisterValueByTarget(target)
	new_value := cpu.regs.a ^ value
	cpu.updateFlags(value, new_value, false, true)
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
}
