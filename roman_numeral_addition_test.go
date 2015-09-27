package main

import (
	"testing"
)

var conversionTests = map[string]int{
	"IX":   9,
	"IIX":  8,
	"XII":  12,
	"VII":  7,
	"IV":   4,
	"":     0,
	"IIMM": 1998,
	"III":  3,
}

func TestParseNum(t *testing.T) {
	for numeral, value := range conversionTests {
		result := parseNum(numeral)
		if result != value {
			t.Errorf("got unexpected value for roman numeral: %s, expected: %d, got %d", numeral, value, result)
		}
	}
}

func TestToRoman(t *testing.T) {
	for numeral, value := range conversionTests {
		result := toRoman(value)
		if result != numeral {
			t.Errorf("got unexpected value for number: %d, expected: %s, got %s", value, numeral, result)
		}
	}
}

func TestAddRoman(t *testing.T) {
	tests := [][]string{
		{"I", "I", "II"},
		{"II", "I", "III"},
		{"IV", "I", "V"},
		{"V", "I", "VI"},
		{"VI", "I", "VII"},
		{"I", "IX", "X"},
		{"X", "IV", "IXV"},
	}

	for test := 0; test < len(tests); test++ {
		result := AddRoman(tests[test][0], tests[test][1])
		if result != tests[test][2] {
			t.Errorf("%s + %s = %s, got: %s", tests[test][0], tests[test][1], tests[test][2], result)
		}
	}
}
