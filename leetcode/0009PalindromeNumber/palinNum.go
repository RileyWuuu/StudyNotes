package main

import "fmt"

func main() {
	result := IsPalindrome(1567652)
	fmt.Println(result)
}
func IsPalindrome(x int) bool {
	var rn int
	tmp := x
	for tmp > 0 {
		rn = rn*10 + tmp%10
		tmp = tmp / 10
	}
	return x == rn
}
