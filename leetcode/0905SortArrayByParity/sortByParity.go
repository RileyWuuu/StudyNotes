package main

import "fmt"

func main() {
	var arr []int
	arr = append(arr, 3, 1, 2, 4)
	fmt.Println(sortArrayByParity(arr))
}

//TimeComplexity:O(n)
func sortArrayByParity(nums []int) []int {
	x := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[x], nums[i] = nums[i], nums[x]
			x++
		}
	}
	return nums
}
