package main

import "fmt"

func main() {
	inputNum := make([]int, 1)
	inputNum = append(inputNum, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(TwoSums(inputNum, 10))
}

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
