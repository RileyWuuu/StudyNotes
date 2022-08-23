package main

import "fmt"

func main() {
	var nums []int
	nums = append(nums, 1, 1, 0, 0, 1, 1, 1, 0, 1)
	fmt.Println(FindMaxConsecutiveOnes(nums))
	fmt.Println(FindMaxForRange(nums))
}

func FindMaxConsecutiveOnes(nums []int) int {
	var result int
	var slice []int
	var max int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			result++
		} else {
			slice = append(slice, result)
			result = 0
		}
		if i == len(nums)-1 {
			slice = append(slice, result)
		}
	}
	max = slice[0]
	for x := 0; x < len(slice); x++ {
		if slice[x] > max {
			max = slice[x]
		}
	}
	return max
}

func FindMaxForRange(nums []int) int {
	res, curr := 0, 0

	for _, n := range nums {
		curr = curr*n + n
		if curr > res {
			res = curr
		}
	}
	return res
}
