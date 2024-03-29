package main

import "fmt"

func main() {
	var nums []int
	nums = append(nums, -2, 0, 10, -19, 4, 6, -8)
	fmt.Println(checkIfExist(nums))
}

func checkIfExist(arr []int) bool {
	intMap := make(map[int]int)
	for i := range arr {
		intMap[arr[i]] = i
	}
	for i := range arr {
		if ii, ok := intMap[arr[i]*2]; ok && ii != i {
			return true
		}
	}
	return false
}

func checkIfExistII(arr []int) bool {
	ln := len(arr)
	for i := 0; i < ln; i++ {
		for x := i + 1; x < ln; x++ {
			if arr[i] == arr[x]*2 || arr[x] == arr[i]*2 {
				return true
			}
		}
	}
	return false
}
