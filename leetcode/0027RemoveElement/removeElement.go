package main

import "fmt"

func main() {
	var arr []int
	arr = append(arr, 3, 2, 2, 3)
	fmt.Println(removeElement(arr, 3))
}

func removeElement(nums []int, val int) int {
	i := 0
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}
	return i
}