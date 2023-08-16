package main

import "fmt"

func main() {
	inputNum := make([]int, 1)
	inputNum = append(inputNum, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(TwoSums(inputNum, 10))
}

//Time complexity:O(n)
func TwoSums(nums []int, target int) []int {
	intMap := make(map[int]int)
	for i, x := range nums {
		if v, ok := intMap[target-x]; ok {
			return []int{v, i}
		}
		intMap[x] = i
	}
	return []int{}

}
//Time complexity:O(n^2)
func TwoSumsII(nums []int, target int) []int {
	result := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for x := i + 1; x < len(nums); x++ {
			if nums[i]+nums[x] == target {
				result[0] = i
				result[1] = x
			}
		}
	}
	return result
}
