package feature

import "fmt"

func TwoSum(nums []int, target int) {
	length := len(nums)
	var newSlice []int
	var signal bool = false
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			tar := nums[i] + nums[j]
			if tar == target {
				newSlice = append(newSlice, i, j)
				signal = true
			}

		}
		if signal {
			break
		}
	}
	fmt.Printf("%v\n", newSlice)
}
