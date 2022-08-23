package main

import (
	"fmt"
	"math"
)

func main() {
	var arr []int
	arr = append(arr, 1, -2147483648, -2147483648)
	fmt.Println(thirdMax(arr))
}

//TimeComplexity:O(n)
func thirdMax(nums []int) int {
	max, second, third := math.MinInt64, math.MinInt64, math.MinInt64
	for _, v := range nums {
		if v == max || v == second || v == third {
			continue
		}
		switch {
		case v > max:
			max, second, third = v, max, second
		case v > second:
			second, third = v, second
		case v > third:
			third = v
		}
	}
	if third == math.MinInt64 {
		return max
	}
	return third
}
