package main

import "fmt"

func main() {
	// var input [][]int
	input := make([][]int, 2, 3)
	input[0] = append(input[0], 1, 2, 3)
	input[1] = append(input[1], 3, 2, 1)
	// MaximumWealth(input)
	fmt.Println(MaximumWealth(input))
}

func MaximumWealth(accounts [][]int) int {
	var result []int
	var tmp, largerNum int
	for i := 0; i < len(accounts); i++ {
		var wealth int
		for _, y := range accounts[i] {
			wealth = wealth + y
		}
		result = append(result, wealth)
	}
	for _, z := range result {
		if z > tmp {
			tmp = z
			largerNum = tmp
		}
	}
	return largerNum
}
