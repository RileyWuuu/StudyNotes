package main

import "fmt"

func main() {
	fmt.Println(NumberOfSteps(15))
}

func NumberOfSteps(num int) int {
	subtractNum := num
	var result int
	for subtractNum > 0 {
		if subtractNum%2 == 0 {
			subtractNum = subtractNum / 2
			result++
		} else {
			subtractNum = subtractNum - 1
			result++
		}
	}
	return result

}
