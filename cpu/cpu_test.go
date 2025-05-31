package cpu

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestFlagRegistersToBit(t *testing.T) {
	type test struct {
		title    string
		flags    FlagsRegister
		expected uint8
	}
	tests := []test{
		{
			title: "all flags",
			flags: FlagsRegister{
				zero:       true,
				substract:  true,
				half_carry: true,
				carry:      true,
			},
			expected: 240,
		},
		{
			// 11000000
			title: "upper flags",
			flags: FlagsRegister{
				zero:       true,
				substract:  true,
				half_carry: false,
				carry:      false,
			},
			expected: 192,
		},
		{
			// 00110000
			title: "lower flags",
			flags: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: true,
				carry:      true,
			},
			expected: 48,
		},
		{
			// 10000000
			title: "zero flag",
			flags: FlagsRegister{
				zero:       true,
				substract:  false,
				half_carry: false,
				carry:      false,
			},
			expected: 128,
		},
		{
			title: "substract flag",
			flags: FlagsRegister{
				zero:       false,
				substract:  true,
				half_carry: false,
				carry:      false,
			},
			expected: 64,
		},
		{
			title: "half_carry flag",
			flags: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: true,
				carry:      false,
			},
			expected: 32,
		},
		{
			title: "carry flag",
			flags: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: false,
				carry:      true,
			},
			expected: 16,
		},
	}
	for _, unit_test := range tests {
		result := unit_test.flags.toBit()
		if result != unit_test.expected {
			t.Errorf("failed : %s expected : %d got : %d", unit_test.title, unit_test.expected, result)
		} else {
			t.Logf("ok: %s", unit_test.title)
		}
	}
}

func TestFlagRegistersToFlagReg(t *testing.T) {
	type test struct {
		title    string
		expected FlagsRegister
		flags    uint8
	}
	tests := []test{
		{
			title: "all flags",
			expected: FlagsRegister{
				zero:       true,
				substract:  true,
				half_carry: true,
				carry:      true,
			},
			flags: 240,
		},
		{
			// 11000000
			title: "upper flags",
			expected: FlagsRegister{
				zero:       true,
				substract:  true,
				half_carry: false,
				carry:      false,
			},
			flags: 192,
		},
		{
			// 00110000
			title: "lower flags",
			expected: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: true,
				carry:      true,
			},
			flags: 48,
		},
		{
			// 10000000
			title: "zero flag",
			expected: FlagsRegister{
				zero:       true,
				substract:  false,
				half_carry: false,
				carry:      false,
			},
			flags: 128,
		},
		{
			title: "substract flag",
			expected: FlagsRegister{
				zero:       false,
				substract:  true,
				half_carry: false,
				carry:      false,
			},
			flags: 64,
		},
		{
			title: "half_carry flag",
			expected: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: true,
				carry:      false,
			},
			flags: 32,
		},
		{
			title: "carry flag",
			expected: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: false,
				carry:      true,
			},
			flags: 16,
		},
	}
	for _, unit_test := range tests {
		result := toFlagReg(unit_test.flags)
		if result != unit_test.expected {
			t.Errorf("failed : %s expected : %v got : %v", unit_test.title, unit_test.expected, result)
		} else {
			t.Logf("ok: %s", unit_test.title)
		}
	}
}
func TestRegistersSetBC(t *testing.T) {
	type test struct {
		title    string
		input    uint16
		expected Registers
	}
	tests := []test{
		{
			// 00110011
			// 0x33
			// 1010011000101101
			title: "checking register boundries",
			input: 42541,
			// 1010011000101101
			// 1010011000000000
			expected: Registers{
				a: 0,
				b: ((0xA62D & 0xFF00) >> 8),
				c: (0xA62D & 0x00FF),
				d: 0,
				e: 0,
				f: FlagsRegister{},
				h: 0,
				l: 0,
			},
		},
		{
			// 00110011
			// 0x33
			// 1010011000101101
			title: "checking register boundries",
			input: 65535,
			// 1010011000101101
			// 1010011000000000
			expected: Registers{
				a: 0,
				b: 0xFF,
				c: 0xFF,
				d: 0,
				e: 0,
				f: FlagsRegister{},
				h: 0,
				l: 0,
			},
		},
	}

	for _, unit_test := range tests {
		var tmp Registers
		tmp.setBC(unit_test.input)
		if tmp != unit_test.expected {
			t.Errorf("failed : %s expected : %v got : %v", unit_test.title, unit_test.expected, tmp)
		} else {
			t.Logf("ok: %s", unit_test.title)
		}
	}
	// 10100110
	// 00101101
}

func TestRegistersGetBC(t *testing.T) {
	type test struct {
		title    string
		expected uint16
		input    Registers
	}
	tests := []test{
		{
			// 00110011
			// 0x33
			// 1010011000101101
			title:    "checking register getter boundries",
			expected: 42541,
			// 1010011000101101
			// 1010011000000000
			input: Registers{
				a: 0,
				b: ((0xA62D & 0xFF00) >> 8),
				c: (0xA62D & 0x00FF),
				d: 0,
				e: 0,
				f: FlagsRegister{},
				h: 0,
				l: 0,
			},
		},
		{
			// 00110011
			// 0x33
			// 1010011000101101
			title:    "checking register getter boundries",
			expected: 65535,
			// 1010011000101101
			// 1010011000000000
			input: Registers{
				a: 0,
				b: 0xFF,
				c: 0xFF,
				d: 0,
				e: 0,
				f: FlagsRegister{},
				h: 0,
				l: 0,
			},
		},
	}

	for _, unit_test := range tests {
		result := unit_test.input.getBC()
		if result != unit_test.expected {
			t.Errorf("failed : %s expected : %v got : %v", unit_test.title, unit_test.expected, result)
		} else {
			t.Logf("ok: %s", unit_test.title)
		}
	}
	// 10100110
	// 00101101
}

func CompareFlagsRegister(left FlagsRegister, right FlagsRegister) (err error) {

	var error_msg strings.Builder

	if left.zero != right.zero {
		error_msg.WriteString(fmt.Sprintf("the zero flags are not the same : %v != %v \n", left.zero, right.zero))
	}
	if left.substract != right.substract {
		error_msg.WriteString(fmt.Sprintf("the substract flags are not the same : %v != %v \n", left.substract, right.substract))
	}
	if left.half_carry != right.half_carry {
		error_msg.WriteString(fmt.Sprintf("the half_carry flags are not the same : %v != %v \n", left.half_carry, right.half_carry))
	}
	if left.carry != right.carry {
		error_msg.WriteString(fmt.Sprintf("the carry flags are not the same : %v != %v \n", left.carry, right.carry))
	}

	if error_msg.Len() > 0 {
		err = errors.New(error_msg.String())
	}

	return err
}

func TestCPUAdd(t *testing.T) {

	type test struct {
		title        string
		sum_result   reg
		overflew     bool
		cpu          CPU
		flags_result FlagsRegister
		target       ArithmeticTarget
	}

	tests := []test{
		{
			title:      "simple addition",
			sum_result: 3,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: false,
				carry:      false,
			},
			cpu: CPU{
				regs: Registers{
					a: reg(2),
					// 00001010
					// 00011001
					c: reg(1),
				},
			},
			target: C,
		},
		{
			title:      "half carry addition",
			sum_result: 35,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: true,
				carry:      false,
			},
			cpu: CPU{
				regs: Registers{
					a: reg(10),
					// 00001010
					// 00011001
					c: reg(25),
				},
			},
			target: C,
		},
		{
			title:      "zero addition",
			sum_result: 0,
			flags_result: FlagsRegister{
				zero:       true,
				substract:  false,
				half_carry: false,
				carry:      false,
			},
			cpu: CPU{
				regs: Registers{
					a: reg(0),
					// 00001010
					// 00011001
					c: reg(0),
				},
			},
			target: C,
		},
		{
			title:      "overflow addition",
			sum_result: 1,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: true,
				carry:      true,
			},
			cpu: CPU{
				regs: Registers{
					a: reg(255),
					// 00001010
					// 00011001
					c: reg(2),
				},
			},
			target: C,
		},
		{
			title:      "overflow addition to zero",
			sum_result: 0,
			flags_result: FlagsRegister{
				zero:       true,
				substract:  false,
				half_carry: true,
				carry:      true,
			},
			cpu: CPU{
				regs: Registers{
					a: reg(255),
					// 00001010
					// 00011001
					c: reg(1),
				},
			},
			target: C,
		},
		{
			title:      "simple addition (A)",
			sum_result: 48,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: true,
				carry:      false,
			},
			cpu: CPU{
				regs: Registers{
					a: reg(24),
					// 00001010
					// 00011001
				},
			},
			target: A,
		},
		{
			title:      "simple addition (B)",
			sum_result: 35,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: true,
				carry:      false,
			},
			cpu: CPU{
				regs: Registers{
					a: reg(24),
					// 00001010
					// 00011001
					b: reg(11),
				},
			},
			target: B,
		},
		{
			title:      "simple addition (D)",
			sum_result: 35,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: true,
				carry:      false,
			},
			cpu: CPU{
				regs: Registers{
					a: reg(24),
					// 00001010
					// 00011001
					d: reg(11),
				},
			},
			target: D,
		},
		{
			title:      "simple addition (E)",
			sum_result: 35,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: true,
				carry:      false,
			},
			cpu: CPU{
				regs: Registers{
					a: reg(24),
					// 00001010
					// 00011001
					e: reg(11),
				},
			},
			target: E,
		},
		{
			title:      "simple addition (H)",
			sum_result: 35,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: true,
				carry:      false,
			},
			cpu: CPU{
				regs: Registers{
					a: reg(24),
					// 00001010
					// 00011001
					h: reg(11),
				},
			},
			target: H,
		},
		{
			title:      "simple addition (L)",
			sum_result: 35,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: true,
				carry:      false,
			},
			cpu: CPU{
				regs: Registers{
					a: reg(24),
					// 00001010
					// 00011001
					l: reg(11),
				},
			},
			target: L,
		},
	}

	for _, unit_test := range tests {
		// result :=
		unit_test.cpu.execute(ADD, unit_test.target)
		success := true
		if unit_test.cpu.regs.a != unit_test.sum_result {
			t.Errorf(
				"failed : %s expected : %v got : %v",
				unit_test.title,
				unit_test.sum_result,
				unit_test.cpu.regs.a,
			)
			success = false
		}

		flags_comp := CompareFlagsRegister(unit_test.cpu.regs.f, unit_test.flags_result)

		if flags_comp != nil {
			t.Errorf(
				"failed : %s expected : %v got : %v\n",
				unit_test.title,
				unit_test.flags_result,
				unit_test.cpu.regs.f,
			)
			t.Error(flags_comp.Error())
			success = false
		}

		if success {
			t.Logf("ok: %s", unit_test.title)
		}

		// if result != unit_test.expected {
		// 	t.Errorf("failed : %s expected : %v got : %v", unit_test.title, unit_test.expected, result)
		// } else {
		// 	t.Logf("ok: %s", unit_test.title)
		// }
	}
	// 10100110
	// 00101101
}

func TestCPUADC(t *testing.T) {

	type test struct {
		title        string
		sum_result   reg
		overflew     bool
		cpu          CPU
		flags_result FlagsRegister
		target       ArithmeticTarget
	}

	tests := []test{
		{
			title:      "overflow and then adding the carry",
			sum_result: 2,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: true,
				carry:      true,
			},
			cpu: CPU{
				regs: Registers{
					a: reg(255),
					// 00001010
					// 00011001
					c: reg(2),
				},
			},
			target: C,
		},
		{
			title:      "overflow and then adding the carry",
			sum_result: 1,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: true,
				carry:      true,
			},
			cpu: CPU{
				regs: Registers{
					a: reg(255),
					// 00001010
					// 00011001
					c: reg(1),
				},
			},
			target: C,
		},
	}

	for _, unit_test := range tests {
		// t :=
		unit_test.cpu.execute(ADC, unit_test.target)
		success := true
		if unit_test.cpu.regs.a != unit_test.sum_result {
			t.Errorf(
				"failed : %s expected : %v got : %v",
				unit_test.title,
				unit_test.sum_result,
				unit_test.cpu.regs.a,
			)
			success = false
		}

		flags_comp := CompareFlagsRegister(unit_test.cpu.regs.f, unit_test.flags_result)

		if flags_comp != nil {
			t.Errorf(
				"failed : %s expected : %v got : %v\n",
				unit_test.title,
				unit_test.flags_result,
				unit_test.cpu.regs.f,
			)
			t.Error(flags_comp.Error())
			success = false
		}

		if success {
			t.Logf("ok: %s", unit_test.title)
		}

		// if result != unit_test.expected {
		// 	t.Errorf("failed : %s expected : %v got : %v", unit_test.title, unit_test.expected, result)
		// } else {
		// 	t.Logf("ok: %s", unit_test.title)
		// }
	}
	// 10100110
	// 00101101
}

func TestCPUSUB(t *testing.T) {

	type test struct {
		instruction  Instruction
		title        string
		sub_result   reg
		overflew     bool
		cpu          CPU
		flags_result FlagsRegister
		target       ArithmeticTarget
	}

	tests := []test{
		{
			instruction: SUB,
			title:       "simple substraction",
			sub_result:  253,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  true,
				half_carry: true,
				carry:      false,
			},
			cpu: CPU{
				regs: Registers{
					a: reg(255),
					// 00001010
					// 00011001
					c: reg(2),
				},
			},
			target: C,
		},
		{
			instruction: SBC,
			title:       "simple substraction with carry",
			sub_result:  250,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  true,
				half_carry: false,
				carry:      true,
			},
			cpu: CPU{
				regs: Registers{
					a: reg(0),
					// 00001010
					// 00011001
					c: reg(5),
				},
			},
			target: C,
		},
	}

	for _, unit_test := range tests {
		// t :=
		unit_test.cpu.execute(unit_test.instruction, unit_test.target)
		success := true
		if unit_test.cpu.regs.a != unit_test.sub_result {
			t.Errorf(
				"failed : %s expected : %v got : %v",
				unit_test.title,
				unit_test.sub_result,
				unit_test.cpu.regs.a,
			)
			success = false
		}

		flags_comp := CompareFlagsRegister(unit_test.cpu.regs.f, unit_test.flags_result)

		if flags_comp != nil {
			t.Errorf(
				"failed : %s expected : %v got : %v\n",
				unit_test.title,
				unit_test.flags_result,
				unit_test.cpu.regs.f,
			)
			t.Error(flags_comp.Error())
			success = false
		}

		if success {
			t.Logf("ok: %s", unit_test.title)
		}
	}
	// 10100110
	// 00101101
}

func TestBitWise(t *testing.T) {

	type test struct {
		instruction  Instruction
		title        string
		result       reg
		overflew     bool
		cpu          CPU
		flags_result FlagsRegister
		target       ArithmeticTarget
	}

	tests := []test{
		{
			instruction: AND,
			title:       "simple and",
			result:      24,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: true,
				carry:      false,
			},
			cpu: CPU{
				regs: Registers{
					//11001 25
					//11000 24
					a: reg(25),
					// 00001010
					// 00011001
					c: reg(24),
				},
			},
			target: C,
		},
		{
			instruction: OR,
			title:       "simple or",
			result:      25,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: true,
				carry:      false,
			},
			cpu: CPU{
				regs: Registers{
					//11001 25
					//11000 24
					a: reg(25),
					// 00001010
					// 00011001
					c: reg(24),
				},
			},
			target: C,
		},
		{
			instruction: XOR,
			title:       "simple xor",
			result:      1,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: true,
				carry:      false,
			},
			cpu: CPU{
				regs: Registers{
					//11001 25
					//11000 24
					//00001 1
					a: reg(25),
					// 00001010
					// 00011001
					c: reg(24),
				},
			},
			target: C,
		},
		{
			instruction: CP,
			title:       "simple cmp",
			result:      25,
			flags_result: FlagsRegister{
				zero:       true,
				substract:  true,
				half_carry: true,
				carry:      false,
			},
			cpu: CPU{
				regs: Registers{
					//11001 25
					//11000 24
					//00001 1
					a: reg(25),
					// 00001010
					// 00011001
					c: reg(25),
				},
			},
			target: C,
		},
		{
			instruction: CP,
			title:       "simple cmp2",
			result:      25,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  true,
				half_carry: false,
				carry:      true,
			},
			cpu: CPU{
				regs: Registers{
					//11001 25
					//11000 24
					//00001 1
					a: reg(25),
					// 00001010
					// 00011001
					c: reg(34),
				},
			},
			target: C,
		},
	}

	for _, unit_test := range tests {
		// t :=
		unit_test.cpu.execute(unit_test.instruction, unit_test.target)
		success := true
		if unit_test.cpu.regs.a != unit_test.result {
			t.Errorf(
				"failed : %s expected : %v got : %v",
				unit_test.title,
				unit_test.result,
				unit_test.cpu.regs.a,
			)
			success = false
		}

		flags_comp := CompareFlagsRegister(unit_test.cpu.regs.f, unit_test.flags_result)

		if flags_comp != nil {
			t.Errorf(
				"failed : %s expected : %v got : %v\n",
				unit_test.title,
				unit_test.flags_result,
				unit_test.cpu.regs.f,
			)
			t.Error(flags_comp.Error())
			success = false
		}

		if success {
			t.Logf("ok: %s", unit_test.title)
		}
	}
	// 10100110
	// 00101101
}
func TestUtilsInstructions(t *testing.T) {
	type test struct {
		instruction Instruction
		title       string
		expected    reg
		result      *reg
		overflew    bool
		// cpu          CPU
		flags_result FlagsRegister
		target       ArithmeticTarget
	}

	cpu := CPU{
		regs: Registers{
			a: 7,
			b: 0,
			c: 0,
			d: 0,
			e: 0,
			h: 0,
			l: 4,
		},
	}

	tests := []test{
		{
			instruction: INC,
			title:       "simple inc",
			expected:    1,
			result:      &cpu.regs.c,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: false,
				carry:      false,
			},
			target: C,
		},
		{
			instruction: DEC,
			title:       "simple dec",
			expected:    0,
			result:      &cpu.regs.c,
			flags_result: FlagsRegister{
				zero:       true,
				substract:  true,
				half_carry: false,
				carry:      false,
			},
			target: C,
		},
		{
			instruction: INC,
			title:       "simple inc (L)",
			expected:    5,
			result:      &cpu.regs.l,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: false,
				carry:      false,
			},
			target: L,
		},
		{
			instruction: DEC,
			title:       "simple dec (A)",
			expected:    6,
			result:      &cpu.regs.a,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  true,
				half_carry: false,
				carry:      false,
			},
			target: A,
		},
		{
			instruction: CCF,
			title:       "simple cff on",
			expected:    6,
			result:      &cpu.regs.a,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: false,
				carry:      true,
			},
		},
		{
			instruction: CCF,
			title:       "simple cff off",
			expected:    6,
			result:      &cpu.regs.a,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: false,
				carry:      false,
			},
		},
		{
			instruction: SCF,
			title:       "simple scf",
			expected:    6,
			result:      &cpu.regs.a,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: false,
				carry:      true,
			},
		},
		{
			instruction: SCF,
			title:       "simple scf 2",
			expected:    6,
			result:      &cpu.regs.a,
			flags_result: FlagsRegister{
				zero:       false,
				substract:  false,
				half_carry: false,
				carry:      true,
			},
		},
	}

	for _, unit_test := range tests {
		// t :=
		cpu.execute(unit_test.instruction, unit_test.target)
		success := true
		if *unit_test.result != unit_test.expected {
			t.Errorf(
				"failed : %s expected : %v got : %v",
				unit_test.title,
				unit_test.expected,
				*unit_test.result,
			)
			t.Log(cpu)
			success = false
		}

		flags_comp := CompareFlagsRegister(cpu.regs.f, unit_test.flags_result)

		if flags_comp != nil {
			t.Errorf(
				"failed : %s expected : %v got : %v\n",
				unit_test.title,
				unit_test.flags_result,
				cpu.regs.f,
			)
			t.Error(flags_comp.Error())
			success = false
		}

		if success {
			t.Logf("ok: %s", unit_test.title)
		}
	}
	// 10100110
	// 00101101
}

type test struct {
	instruction Instruction
	title       string
	expected    reg
	result      *reg
	regs        Registers
	overflew    bool
	// cpu          CPU
	flags_result FlagsRegister
	target       ArithmeticTarget
}

func (t *test) updateRegisters(cpu *CPU) {
	cpu.regs = t.regs
}

func TestRegisterRotation(t *testing.T) {

	cpu := CPU{}
	// 101001100
	// 10100110
	tests := []test{
		{
			// 01010100
			// 01010100
			regs: Registers{
				d: 84,
				f: FlagsRegister{
					carry: true,
				},
			},
			instruction: RL,
			title:       "simple rl d",
			expected:    169,
			result:      &cpu.regs.d,
			flags_result: FlagsRegister{
				carry: false,
			},
			target: D,
		},
		{
			// 01010100
			// 01010100
			regs: Registers{
				a: 84,
				b: 0,
				c: 0,
				d: 0,
				e: 0,
				h: 0,
				l: 0,
				f: FlagsRegister{
					carry: true,
				},
			},
			instruction: RLA,
			title:       "simple rla",
			expected:    169,
			result:      &cpu.regs.a,
			flags_result: FlagsRegister{
				carry: false,
			},
			target: A,
		},
		{
			regs: Registers{
				a: 76,
				b: 0,
				c: 0,
				d: 0,
				e: 0,
				h: 0,
				l: 0,
				f: FlagsRegister{
					carry: true,
				},
			},
			instruction: RRA,
			title:       "simple rra",
			expected:    166,
			result:      &cpu.regs.a,
			flags_result: FlagsRegister{
				carry: false,
			},
			target: A,
		},
		{
			// 11010100
			// 10101001
			regs: Registers{
				a: 0,
				b: 0,
				c: 212,
				d: 0,
				e: 0,
				h: 0,
				l: 0,
				f: FlagsRegister{
					carry: false,
				},
			},
			instruction: RLC,
			title:       "simple rlc c",
			expected:    169,
			result:      &cpu.regs.c,
			flags_result: FlagsRegister{
				carry: true,
			},
			target: C,
		},
		{
			// 11010100
			// 10101001
			regs: Registers{
				a: 212,
				b: 0,
				c: 0,
				d: 0,
				e: 0,
				h: 0,
				l: 0,
				f: FlagsRegister{
					carry: false,
				},
			},
			instruction: RLCA,
			title:       "simple rlca",
			expected:    169,
			result:      &cpu.regs.a,
			flags_result: FlagsRegister{
				carry: true,
			},
			target: A,
		},
		{
			// c:0
			// d:01010100
			// d:10101010
			regs: Registers{
				a: 84,
				f: FlagsRegister{
					carry: true,
				},
			},
			instruction: RRA,
			title:       "simple rra",
			expected:    170,
			result:      &cpu.regs.a,
			flags_result: FlagsRegister{
				carry: false,
			},
			target: A,
		},
		{
			// c:0
			// d:01010100
			// d:10101010
			regs: Registers{
				d: 84,
				f: FlagsRegister{
					carry: true,
				},
			},
			instruction: RR,
			title:       "simple rr d",
			expected:    170,
			result:      &cpu.regs.d,
			flags_result: FlagsRegister{
				carry: false,
			},
			target: D,
		},
		{
			// c:0
			// d:01010101
			// d:10101010
			// c:1
			regs: Registers{
				d: 0b01010101,
				f: FlagsRegister{
					carry: false,
				},
			},
			instruction: RRC,
			title:       "simple rrc d",
			expected:    0b10101010,
			result:      &cpu.regs.d,
			flags_result: FlagsRegister{
				carry: true,
			},
			target: D,
		},
		{
			// c:0
			// d:01010101
			// d:10101010
			// c:1
			regs: Registers{
				a: 0b01010101,
				f: FlagsRegister{
					carry: false,
				},
			},
			instruction: RRCA,
			title:       "simple rrca",
			expected:    0b10101010,
			result:      &cpu.regs.a,
			flags_result: FlagsRegister{
				carry: true,
			},
			target: A,
		},
		{
			// c:0
			// a:10010101
			// a:00101010
			// c:1
			regs: Registers{
				a: 0b10010101,
				f: FlagsRegister{
					carry: false,
				},
			},
			instruction: SLL,
			title:       "simple sll a",
			expected:    0b00101010,
			result:      &cpu.regs.a,
			flags_result: FlagsRegister{
				carry: true,
			},
			target: A,
		},
		{
			// make somethign to parse this
			// input for the tests and make
			// a test based on that
			// c:0
			// a:10010101
			// a:01001010
			// c:1
			regs: Registers{
				a: 0b10010101,
				f: FlagsRegister{
					carry: false,
				},
			},
			instruction: SRL,
			title:       "simple srl a",
			expected:    0b01001010,
			result:      &cpu.regs.a,
			flags_result: FlagsRegister{
				carry: true,
			},
			target: A,
		},
		{
			// make somethign to parse this
			// input for the tests and make
			// a test based on that
			// c:0
			// a:10010101
			// a:11001010
			// c:1
			regs: Registers{
				a: 0b10010101,
				f: FlagsRegister{
					carry: false,
				},
			},
			instruction: SRA,
			title:       "simple sra a",
			expected:    0b11001010,
			result:      &cpu.regs.a,
			flags_result: FlagsRegister{
				carry: true,
			},
			target: A,
		},
		{
			// c:0
			// a:00010101
			// a:00001010
			// c:1
			regs: Registers{
				a: 0b00010101,
				f: FlagsRegister{
					carry: false,
				},
			},
			instruction: SRA,
			title:       "simple sra a ari shift but with zero in the lowest",
			expected:    0b00001010,
			result:      &cpu.regs.a,
			flags_result: FlagsRegister{
				carry: true,
			},
			target: A,
		},
		{
			// make somethign to parse this
			// input for the tests and make
			// a test based on that
			// c:0
			// a:10010101
			// a:00101011
			// c:1
			regs: Registers{
				a: 0b10010101,
				f: FlagsRegister{
					carry: false,
				},
			},
			instruction: SLA,
			title:       "simple sla",
			expected:    0b00101011,
			result:      &cpu.regs.a,
			flags_result: FlagsRegister{
				carry: true,
			},
			target: A,
		},
		{
			// c:0
			// a:00010100
			// a:00101000
			// c:1
			regs: Registers{
				a: 0b10010100,
				f: FlagsRegister{
					carry: false,
				},
			},
			instruction: SLA,
			title:       "simple sla a ari shift but with zero in the highest",
			expected:    0b00101000,
			result:      &cpu.regs.a,
			flags_result: FlagsRegister{
				carry: true,
			},
			target: A,
		},
	}

	for _, unit_test := range tests {
		// t :=
		unit_test.updateRegisters(&cpu)
		cpu.execute(unit_test.instruction, unit_test.target)
		success := true
		if *unit_test.result != unit_test.expected {
			t.Errorf(
				"failed : %s expected : %b got : %b",
				unit_test.title,
				unit_test.expected,
				*unit_test.result,
			)
			t.Log(cpu)
			success = false
		}

		flags_comp := CompareFlagsRegister(cpu.regs.f, unit_test.flags_result)

		if flags_comp != nil {
			t.Errorf(
				"failed : %s expected : %v got : %v\n",
				unit_test.title,
				unit_test.flags_result,
				cpu.regs.f,
			)
			t.Error(flags_comp.Error())
			success = false
		}

		if success {
			t.Logf("ok: %s", unit_test.title)
		}
	}
	// 10100110
	// 00101101
}

func TestBitManipulation(t *testing.T) {

	cpu := CPU{}
	// 101001100
	// 10100110
	tests := []test{
		{
			// 01010100
			// 01010100
			regs: Registers{
				a: 0b00000000,
			},
			instruction: CPL,
			title:       "simple cpl",
			expected:    0b11111111,
			result:      &cpu.regs.a,
			flags_result: FlagsRegister{
				substract:  true,
				half_carry: true,
			},
			target: A,
		},
		{
			// 01010100
			// 01010100
			regs: Registers{
				a: 0b10100101,
			},
			instruction: CPL,
			title:       "simple cpl complex input",
			expected:    0b01011010,
			result:      &cpu.regs.a,
			flags_result: FlagsRegister{
				substract:  true,
				half_carry: true,
			},
			target: A,
		},
	}

	for _, unit_test := range tests {
		// t :=
		unit_test.updateRegisters(&cpu)
		cpu.execute(unit_test.instruction, unit_test.target)
		success := true
		if *unit_test.result != unit_test.expected {
			t.Errorf(
				"failed : %s expected : %b got : %b",
				unit_test.title,
				unit_test.expected,
				*unit_test.result,
			)
			t.Log(cpu)
			success = false
		}

		flags_comp := CompareFlagsRegister(cpu.regs.f, unit_test.flags_result)

		if flags_comp != nil {
			t.Errorf(
				"failed : %s expected : %v got : %v\n",
				unit_test.title,
				unit_test.flags_result,
				cpu.regs.f,
			)
			t.Error(flags_comp.Error())
			success = false
		}

		if success {
			t.Logf("ok: %s", unit_test.title)
		}
	}
	// 10100110
	// 00101101
}

func TestBitBit(t *testing.T) {

	cpu := CPU{}
	// in the tests we will store the argument for the bit to check in the
	// d register
	tests := []test{
		{
			// 01010100
			// 01010100
			regs: Registers{
				a: 0b00000001,
				d: 0,
			},

			instruction: BIT,
			title:       "bit a",
			result:      &cpu.regs.a,
			flags_result: FlagsRegister{
				half_carry: true,
				zero:       false,
			},
			target: A,
		},
	}

	tests[0].updateRegisters(&cpu)
	for _, unit_test := range tests {
		// t :=
		for b := range 8 {
			t.Logf("reg a: %08b", cpu.regs.a)
			cpu.execute(unit_test.instruction, unit_test.target, uint8(b))
			success := true

			flags_comp := CompareFlagsRegister(cpu.regs.f, unit_test.flags_result)
			if flags_comp != nil {
				t.Errorf(
					"failed : %s expected : %v got : %v\n",
					unit_test.title,
					unit_test.flags_result,
					cpu.regs.f,
				)
				t.Error(flags_comp.Error())
				success = false
			}

			if success {
				t.Logf("ok: %s %d", unit_test.title, b)
			}
			cpu.regs.a <<= 1
		}

	}
	// 10100110
	// 00101101
}

func TestBitReset(t *testing.T) {

	cpu := CPU{}
	// in the tests we will store the argument for the bit to check in the
	// d register
	tests := []test{
		{
			// 01010100
			// 01010100
			regs: Registers{
				a: 0b11111111,
				d: 0,
			},

			instruction: RESET,
			title:       "reset a",
			result:      &cpu.regs.a,
			target:      A,
		},
	}

	for _, unit_test := range tests {
		unit_test.updateRegisters(&cpu)
		// t :=
		for b := range 8 {
			cpu.execute(unit_test.instruction, unit_test.target, uint8(b))
			success := true

			flags_comp := CompareFlagsRegister(cpu.regs.f, unit_test.flags_result)
			if flags_comp != nil {
				t.Errorf(
					"failed : %s expected : %v got : %v\n",
					unit_test.title,
					unit_test.flags_result,
					cpu.regs.f,
				)
				t.Error(flags_comp.Error())
				success = false
			}

			if success {
				t.Logf("ok: %s %d", unit_test.title, b)
				t.Logf("reg a: %08b", cpu.regs.a)
			}
		}

	}
	// 10100110
	// 00101101
}

func TestBitSet(t *testing.T) {

	cpu := CPU{}
	// in the tests we will store the argument for the bit to check in the
	// d register
	tests := []test{
		{
			// 01010100
			// 01010100
			regs: Registers{
				a: 0b00000000,
				d: 0,
			},

			instruction: SET,
			title:       "reset a",
			result:      &cpu.regs.a,
			target:      A,
		},
	}

	for _, unit_test := range tests {
		unit_test.updateRegisters(&cpu)
		// t :=
		for b := range 8 {
			cpu.execute(unit_test.instruction, unit_test.target, uint8(b))
			success := true

			flags_comp := CompareFlagsRegister(cpu.regs.f, unit_test.flags_result)
			if flags_comp != nil {
				t.Errorf(
					"failed : %s expected : %v got : %v\n",
					unit_test.title,
					unit_test.flags_result,
					cpu.regs.f,
				)
				t.Error(flags_comp.Error())
				success = false
			}

			if success {
				t.Logf("ok: %s %d", unit_test.title, b)
				t.Logf("reg a: %08b", cpu.regs.a)
			}
		}

	}
	// 10100110
	// 00101101
}
