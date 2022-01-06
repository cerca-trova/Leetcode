package feature

func twoSum(nums []int, target int) []int {
	length := len(nums)
	var newSlice []int
	var signal bool
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			tar := nums[i] + nums[j]
			if tar == target {
				newSlice = append(newSlice, i, j)
				signal = true
			}

		}
		if signal == true {
			break
		}
	}
	return newSlice
}
