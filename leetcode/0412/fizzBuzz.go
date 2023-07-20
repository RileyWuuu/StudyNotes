package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FizzBuzz(9))
}

func FizzBuzz(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		if i%5 == 0 && i%3 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}
