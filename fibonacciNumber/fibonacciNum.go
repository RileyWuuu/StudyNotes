package main

import "fmt"

func main() {
	x := fibonacciNum(3)
	fmt.Println(x)
}
func fibonacciNum(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacciNum(n-1) + (n - 2)
}
