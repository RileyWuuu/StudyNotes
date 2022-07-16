package main

import "fmt"

func main() {
	x := FibonacciNum(1)
	fmt.Println(x)
}
func FibonacciNum(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return FibonacciNum(n-1) + FibonacciNum(n-2)
}
