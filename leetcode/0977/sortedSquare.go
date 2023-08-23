package main

//TimeComplexity:nlogn
// func sortedSquare(nums []int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		nums[i] = int(math.Pow(float64(nums[i]), 2))
// 	}
// 	sort.Ints(nums)
// 	return nums
// }

//TimeComplexity:O(n)
// math.pow(x,2)=x二次方
// math.abs(float(x))=負數轉正

func sortedSquareII(nums []int) []int {
	for i := range nums {
		nums[i] *= nums[i]
	}
	idx := len(nums) - 1
	head, tail := 0, idx
	result := make([]int, len(nums))
	for idx >= 0 {
		if nums[head] > nums[tail] {
			result[idx] = head
			head++
		} else {
			result[idx] = tail
			tail--
		}
		idx--
	}
	return result
}
