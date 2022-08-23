package main

import (
	"sort"
)

func main() {
	inputInt := make([]int, 1)
	inputInt = append(inputInt, 1, 2, 3, 4, 5, 6, 7)

	SearchInsert(inputInt, 4)
	SortSearch(inputInt,4)

}
func SearchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}
func SortSearch(nums []int,target int)int{
	i:=nums.sort.Search(len(nums)func(i int)bool{
		return nums[i] >= target
	})
	return i
}

