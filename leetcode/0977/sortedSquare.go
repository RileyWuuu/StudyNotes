package main

import (
	"fmt"
	"sort"
)
//TimeComplexity:nlogn
// func sortedSquare(nums []int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		nums[i] = int(math.Pow(float64(nums[i]), 2))
// 	}
// 	sort.Ints(nums)
// 	return nums
// }

//TimeComplexity:O(n)
// math.pow(x,2)=x二次方
// math.abs(float(x))=負數轉正

func sortedSquareII(nums []int)[]int{
	head:=0
	tail:=len(nums)-1
	result := make([]int,len(nums))
	current:=len(nums)-1
	for current>=0{
		if math.Abs(float(nums[head]))>=math.Abs(float(nums[tail])){
			result[current]=int(math.Pow(float(nums[head]),2))
			head++
		}else{
			result[current]=int(math.Pow(float(nums[tail]),2))
			tail--
		}
		current--
	}
	return result
}
