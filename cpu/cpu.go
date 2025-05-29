package cpu

// la funcion de toBit la pude hacer porque sabia que las expressiones finales de los if statements son a lo que terminan evaluandose, por lo que probablemente haya mas cosas que no tenga la menor idea asi que tendre que buscar la referencia de lo que significa en internet o preguntandole a alguient
type Instruction uint8

type CPU struct {
	regs Registers
}

func (cpu *CPU) setZeroFlag(value bool) {
	cpu.regs.f.zero = value
}
