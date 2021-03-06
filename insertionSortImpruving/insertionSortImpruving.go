package main

import (
	"fmt"

	"github.com/dreddsa5dies/algorithm/util"
)

func main() {
	s1 := util.RandomInt() // срез int
	fmt.Printf("Unsorted list:\t%v\n", s1)
	fmt.Println("")
	for i := 1; i < len(s1); i++ {
		value := s1[i]
		for j := i; j != 0 && value < s1[j-1]; s1[j] = value {
			s1[j] = s1[j-1]
			j--
		}
		fmt.Printf("Sorting ...:\t%v\n", s1)
	}
	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", s1)
}
