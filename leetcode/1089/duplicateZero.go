package main

import "fmt"

func main() {
	var arr []int
	arr = append(arr, 1, 0, 9, 4, 0, 7, 0, 3, 6)
	duplicateZeros(arr)
}

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 && i+1 < len(arr) {
			arr = append(arr[:i+1], arr[i:len(arr)-1]...)
			i++
		}
	}
	fmt.Println(arr)
}
