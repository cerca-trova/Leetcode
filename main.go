package main

import (
	"fmt"

	functions "github.com/cerca-trova/Leetcode/feature"
)

func main() {
	var s = []int{1, 2, 3, 4}
	s1 := s[0:3]
	if functions.ValidParentheses("[()]") {
		fmt.Println("Correct!")
	} else {
		fmt.Println("NOT Correct!")
	}
	for i := 0; i < len(s1); i++ {
		if s[i] == 1 {
			i = 2

		}
		fmt.Println(s[i])
		continue
	}

}
