package main

import (
	"fmt"

	constants "github.com/cerca-trova/Leetcode/constant"
	functions "github.com/cerca-trova/Leetcode/feature"
)

func main() {
	if functions.ValidParentheses(constants.ValidParenthesesSample3) {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Not Correct!")
	}
}
