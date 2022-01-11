package feature

import "fmt"

var ValidParenthesesAnsiNums = []int{40, 41, 123, 125, 91, 93}

func ValidParentheses(s string) {
	length := len(s)
	if length%2 != 0 || s[0] == 41 || s[0] == 125 || s[0] == 93 {
		fmt.Println(0)
	} else {
		givenParenthesesPairs := length / 2
	}

}
