package main

import (
	"testing"
)

var conversionTests = map[string]int{
    "IX":  9,
    "IIX": 8,
    "XII": 12,
    "IIIX": 7,
    "IVX": 4,
    "VVX": 0,
    "VIX": 4,
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
