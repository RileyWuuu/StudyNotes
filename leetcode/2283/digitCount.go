package main

import (
	"fmt"
)

func main() {
	fmt.Println(DigitCount("1210"))
}

func DigitCount(num string) bool {

	var a [10]int
	var b [10]int
	for i := 0; i < len(num); i++ {
		a[num[i]-'0']++
		b[i] = int(num[i] - '0')
	}
	return a == b
}
