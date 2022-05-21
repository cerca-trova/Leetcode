package main

import (
	"fmt"

	functions "github.com/cerca-trova/Leetcode/feature"
)

func main() {
	if functions.ValidParentheses("[()][]") {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Not Correct!")
	}
}
