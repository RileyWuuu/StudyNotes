package main

import "fmt"

func main() {
	x := climbStairs(40)
	fmt.Println(x)
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-1) + (n - 2)
}
