package p0724

//TimeComplexity: O(n)
func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	count := make([]int, len(nums))
	count[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		count[i] = count[i-1] + nums[i]
	}
	if count[len(nums)-1]-count[0] == 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if count[i-1] == count[len(nums)-1]-count[i] {
			return i
		}
	}
	return -1
}
