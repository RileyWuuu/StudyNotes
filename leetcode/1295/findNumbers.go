package main

import "fmt"

func main() {
	var data []int
	data = append(data, 123, 45, 6912, 33, 2, 4, 758264)
	fmt.Println(FindNumbers(data))
}

func FindNumbers(nums []int) int {
	var digit []int
	var result int

	for _, x := range nums {
		var count int
		for x > 0 {
			x = x / 10
			count++
		}
		digit = append(digit, count)
	}
	for i := 0; i < len(digit); i++ {
		if digit[i]%2 == 0 {
			result++
		}
	}
	return result
}

func FindNumbersSimple(nums []int)int{
	var result int
	for x :=range nums{
		i:=0
		for nums[x] >0{
			i++
			nums[x] = nums[x]/10
		}
		if i %2==0{
			result++
		}
	}
	return result
}
