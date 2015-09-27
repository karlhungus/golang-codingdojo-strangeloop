package main

import (
	"fmt"
)

var numberMaps = map[string]int{
	"M": 1000,
	"C": 100,
	"X": 10,
	"V": 5,
	"I": 1,
}

var numberArray = []map[string]int{
	{"M": 1000},
	{"C": 100},
	{"X": 10},
	{"V": 5},
	{"I": 1},
}

func main() {
	fmt.Printf("foo")
	input := "I+I"
	fmt.Printf("input: %s\n", input)
	fmt.Printf("output: %s\n", AddRoman("I", "I"))
}

func AddRoman(s1 string, s2 string) string {
	return toRoman(parseNum(s1) + parseNum(s2))
}

func parseNum(s string) int {
	var count = 0
	var max = 0
	var array = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		num := numberMaps[string(s[i])]
		array[i] = num
		if max < num {
			max = num
		}
	}
	max_hit := false
	for j := 0; j < len(array); j++ {
		if array[j] == max {
			max_hit = true
		}
		if !max_hit {
			count -= array[j]
		} else {
			count += array[j]
		}
	}
	return count
}

func toRoman(input int) string {
	var store string
	remainder := input
	for i := 0; i < len(numberArray); i++ {
		for numeral, value := range numberArray[i] {
			test := value
			if value > 5 {
				test = value - 2
			} else if value == 5 {
				test = value - 1
			}
			for remainder >= test {
				remainder -= value
				store += numeral
			}
			for remainder < 0 && remainder >= -2 {
				remainder += 1
				store = "I" + store
			}
		}
	}
	return store
}
