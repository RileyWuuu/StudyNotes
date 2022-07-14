package main

import "fmt"

func main() {
	inputInt := make([]int, 1)
	inputInt = append(inputInt, 1, 2, 3, 4, 5, 6, 7)

	result := SearchInsert(inputInt, 4)
	fmt.Println("result=", result)

}
func SearchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			fmt.Println("i=", i)
			return i
		}
	}
	return len(nums)
}
