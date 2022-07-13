package main

import "fmt"

func uniqueNames(a, b []string) {
	var result []string
	for x := range a {
		result = append(result, a[x])
	}
	for y := range b {
		result = append(result, b[y])
	}
	x := removeDuplicate(result)
	fmt.Println(x)

}
func removeDuplicate(slice []string) []string {
	result := len(slice)
	if result <= 1 {
		return slice[0:result]
	}
	j := 1
	for i := 1; i < result; i++ {
		if slice[i] != slice[i-1] {
			slice[j] = slice[i]
			j++
		}
	}
	return slice[0:j]
}
func main() {
	// should print Ava, Emma, Olivia, Sophia
	uniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"})
}
