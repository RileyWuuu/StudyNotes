package main

func ProductExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	multiply := 1
	for i := 0; i < len(nums); i++ {
		// number itself doesn't count
		res[i] = multiply
		multiply *= nums[i]
	}
	multiply = 1
	for i := len(res) - 1; i >= 0; i-- {
		res[i] *= multiply
		multiply *= nums[i]
	}
	return res
}
