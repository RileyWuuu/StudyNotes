package main

import "fmt"

func main() {
	var arr []int
	arr = append(arr, 1, 5, 7, 6, 1, 3, 3)
	fmt.Println(countSort(arr))
}

func countSort(arr []int) []int {
	count := make([]int, len(arr)+1)
	for _, x := range arr {
		count[x]++
	}
	for i, k := 1, 0; i < len(count); i++ {
		for z := count[i]; z > 0; z-- {
			arr[k] = i
			k++
		}
	}
	return arr
}
