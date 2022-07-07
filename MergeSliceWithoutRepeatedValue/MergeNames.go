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
	removeDuplicate(result)
}
func removeDuplicate(slice []string) []string {
	result := []string{}
	tempMap := map[string]byte{}
	for _, z := range slice {
		length := len(tempMap)
		tempMap[z] = 0
		if len(tempMap) != length {
			result = append(result, z)
		}
	}
	fmt.Println(result)
	return result
}
func main() {
	// should print Ava, Emma, Olivia, Sophia
	uniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"})
}
