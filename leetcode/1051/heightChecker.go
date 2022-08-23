package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr []int
	arr = append(arr, 1, 1, 4, 2, 1, 3)
	fmt.Println(heightChecker(arr))
	fmt.Println(heightCheckerII(arr))
}

//Time Complexity: nlogn
func heightChecker(heights []int) int {
	var sliceInOrder []int
	sliceInOrder = append(sliceInOrder, heights...)
	sort.Ints(sliceInOrder)
	count := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != sliceInOrder[i] {
			count++
		}
	}
	return count
}

//Time Complexity: O(n)
func heightCheckerII(heights []int) int {
	counts := make([]int, 101)
	for _, x := range heights {
		counts[x]++
	}
	ans := 0
	for i, z := 1, 0; i < len(counts); i++ {
		for k := counts[i]; k > 0; k-- {
			if i != heights[z] {
				ans++
			}
			z++
		}
	}
	return ans
}
