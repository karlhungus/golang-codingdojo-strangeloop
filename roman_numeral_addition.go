package main

import "fmt"

var numberMaps = map[string]int {
  "I": 1,
  "V": 5,
  "X": 10,
  "C": 100,
}

func main() {
  fmt.Printf("foo")
  input := "I+I"
  fmt.Printf("input: %s\n", input)
  fmt.Printf("output: %d\n", AddRoman("I", "I"))
  fmt.Printf(" IX(9): %d\n",parseNum("IX"))
  fmt.Printf(" IIX(8): %d\n",parseNum("IIX"))
  fmt.Printf(" XII(12): %d\n",parseNum("XII"))
  fmt.Printf(" IIX(7): %d\n",parseNum("IIIX"))
  fmt.Printf(" IVX(4): %d\n",parseNum("IVX"))
  fmt.Printf(" VVX(0): %d\n",parseNum("IVX"))
}

func AddRoman(s1 string, s2 string) int {
  return parseNum(s1) + parseNum(s2)
}

func parseNum(s string) int {
  var count = 0
  var max = 0
  var array = make([]int,len(s))
  for i := 0; i < len(s); i++ {
    num := numberMaps[string(s[i])]
    //fmt.Printf("%q\n", s[i])
    //fmt.Printf("%d\n", )
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
