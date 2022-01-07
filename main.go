package main

import (
	"fmt"

	variables "github.com/cerca-trova/Leetcode/constant"
	functions "github.com/cerca-trova/Leetcode/feature"
)

func main() {
	functions.TwoSum(variables.TestSlice2, variables.Target2)
	var j int = 2
	var i int = 10 >> j
	d := calDigits(i)
	fmt.Println("i is \n", i)
	fmt.Println("d is \n", d)

}
func calDigits(x int) int {
	var digits int = 1
	for i := x; i >= 10; i /= 10 {
		digits++
	}
	return digits
}
