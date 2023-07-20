package p0905

//TimeComplexity:O(n)
func SortArrayByParity(nums []int) []int {
	x := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[x], nums[i] = nums[i], nums[x]
			x++
		}
	}
	return nums
}
