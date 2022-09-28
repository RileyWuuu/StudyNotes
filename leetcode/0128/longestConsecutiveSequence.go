package p0128

//TimeComplexity O(n)
func longestConsecutive(nums []int) int {
	result := 0
	m := make(map[int]bool)
	for _, x := range nums {
		m[x] = true
	}
	for x := range m {
		if m[x-1] {
			continue
		}
		count := x
		for m[count+1] {
			count++
		}
		result = max(result, count-x+1)
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
