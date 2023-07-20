package main

import "fmt"

func main() {
	var input []int
	input = append(input, 1, 3, 5, 7, 9)
	fmt.Println(RunningSum(input))
}

func RunningSum(nums []int) []int {

	var result []int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result = append(result, nums[i])
		} else {
			result = append(result, nums[i-1]+nums[i])
			nums[i] = nums[i-1] + nums[i]
		}
	}
	fmt.Println("RESULT:", result)
	return result
}
