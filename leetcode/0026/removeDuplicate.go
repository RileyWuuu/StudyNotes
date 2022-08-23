package main

import "fmt"

func main() {
	var nums []int
	nums = append(nums, 0, 0, 1, 1, 1, 2, 3, 3, 3)
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	x := 0
	if len(nums) <= 1 {
		return len(nums)
	}
	for i := 1; i < len(nums); i++ {
		if nums[x] != nums[i] {
			x++
			nums[x] = nums[i]
		}
	}
	return x + 1
}
