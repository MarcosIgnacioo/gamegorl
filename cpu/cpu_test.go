package cpu

import (
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
				f: 0,
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
				f: 0,
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
				f: 0,
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
				f: 0,
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
