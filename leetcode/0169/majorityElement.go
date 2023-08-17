package p0128

func majorityElement(nums []int) int {
	maxNum, resp := 0, 0
	elements := make(map[int]int)
	for _, val := range nums {
		elements[val]++
	}
	for index, val := range elements {
		if val > maxNum {
			maxNum = val
			resp = index
		}
	}
	return resp
}
