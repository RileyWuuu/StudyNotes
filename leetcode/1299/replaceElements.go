package main

import "fmt"

func main() {
	var arr []int
	arr = append(arr, 17, 18, 5, 4, 6, 1)
	fmt.Println(replaceElements(arr))
}
func replaceElements(arr []int) []int {
	max := -1
	ln := len(arr)
	for i := ln - 1; i >= 0; i-- {
		tmp := max
		if arr[i] > max {
			max = arr[i]
		}
		arr[i] = tmp
	}
	return arr

}
