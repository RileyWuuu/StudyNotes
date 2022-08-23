package p0747

//TimeComplexity:O(n)
func dominantIndex(nums []int) int {
	max, idx := 0, 0
	for i, x := range nums {
		if max < x {
			max, idx = x, i
		}
	}
	for _, x := range nums {
		if x != max && x*2 > max {
			return -1
		}
	}
	return idx
}
