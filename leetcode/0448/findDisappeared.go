package main

import "fmt"

func main() {
	var arr []int
	arr = append(arr, 4, 3, 2, 7, 8, 2, 3, 1)
	fmt.Println(findDisappearedNumbers(arr))

}

//Time Complexity: O(n)
func findDisappearedNumbers(nums []int) []int {
	result := []int{}
	arr := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		arr[nums[i]-1] += 1
	}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			result = append(result, i+1)
		}
	}
	return result
}
