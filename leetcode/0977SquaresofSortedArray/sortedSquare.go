package main

import (
	"fmt"
	"sort"
)

func main() {
	var data []int
	data = append(data, -11, -3, 0, 5, 19)
	fmt.Println(SortedSquare(data))
}

func SortedSquare(nums []int) []int {
	for x := range nums {
		nums[x] = nums[x] * nums[x]
	}
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})

	return nums
}
