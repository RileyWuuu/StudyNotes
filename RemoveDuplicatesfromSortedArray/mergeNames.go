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
	x := RemoveDuplicate(result)
	fmt.Println(x)
}

func RemoveDuplicate(slice []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range slice {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}
func main() {
	// should print Ava, Emma, Olivia, Sophia
	uniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Sophia", "Olivia", "Emma"})
}
