package main

import (
	"sort"
)

func SearchInsert(nums []int, target int) int {
	i := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	return i
}
