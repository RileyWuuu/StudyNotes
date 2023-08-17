package main

func containsDuplicate(nums []int) bool {
	valMap := make(map[int]bool)
	for _, val := range nums {
		if _, ok := valMap[val]; ok {
			return true
		}
		valMap[val] = true
	}
	return false
}
