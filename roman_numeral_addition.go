package main

import "fmt"

var numberMaps = map[string]int{
  "M" : 1000,
	"C": 100,
	"X": 10,
	"V": 5,
	"I": 1,
}

var numberArray = []map[string]int {
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
	fmt.Printf("output: %d\n", AddRoman("I", "I"))
}

func AddRoman(s1 string, s2 string) int {
	return parseNum(s1) + parseNum(s2)
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
 fmt.Printf("input:%d\n===============\n", input)
  var store string
  remainder := input
  for i := 0; i < len(numberArray); i++ {
    for numeral, value := range numberArray[i] {
      for remainder >= value {
        remainder -= value
        store += numeral
      }
      fmt.Printf("remainder:%d, numeral:%s, value:%d, store:%s\n", remainder, numeral, value, store)
    }
  }
  return store
}

