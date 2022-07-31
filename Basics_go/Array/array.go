package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := [10]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	arr2 := [10]int{1, 2, 3, 4, 5}
	fmt.Println("writeWithIndex:", WriteWithIndex(arr, "STOP"))
	fmt.Println("writeWithloop:", WriteWithloop(arr, "123"))
	fmt.Println("readWithloop:", ReadWithloop(arr))
	fmt.Println("readWithForRange:", ReadWithForRange(arr))
	fmt.Println("shiftElement:", insertFistIndex(arr2))
	fmt.Println("insertAnyIndex:", insertAnyIndex(arr2))
	fmt.Println("deleteFirstIndex:", deleteFirstIndex(arr2))
	fmt.Println("deleteAnyIndex:", deleteAnyIndex(arr2))
}

func WriteWithIndex(data [10]string, replace string) [10]string {

	data[2] = replace
	fmt.Println(data[2])
	return data
}

func WriteWithloop(data [10]string, x string) [10]string {
	for i := 0; i < len(data); i++ {
		input := x + strconv.Itoa((i+1)*(i+1))
		data[i] = input
	}
	return data
}

func ReadWithloop(data [10]string) [10]string {
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
	return data
}

func ReadWithForRange(data [10]string) [10]string {
	for _, i := range data {
		fmt.Println(i)
	}
	return data
}
func insertFistIndex(data [10]int) [10]int {

	for i := 5; i >= 0; i-- {
		data[i+1] = data[i]
	}
	data[0] = 777
	return data
}
func insertAnyIndex(data [10]int) [10]int {
	for i := 6; i >= 3; i-- {
		data[i+1] = data[i]
	}
	data[3] = 777
	return data
}

func deleteFirstIndex(data [10]int) [10]int {
	for i := 1; i < len(data); i++ {
		data[i-1] = data[i]
	}
	return data
}

func deleteAnyIndex(data [10]int) [10]int {
	for i := 3; i < len(data); i++ {
		data[i-1] = data[i]
	}
	return data
}


