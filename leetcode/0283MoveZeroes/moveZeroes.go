package main

func main() {
	var arr []int
	arr = append(arr, 0, 1, 0, 3, 12)
	moveZeroes(arr)
}

func moveZeroes(nums []int) {
	length := len(nums)
	x := 0
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[x], nums[i] = nums[i], nums[x]
			x++
		}
	}
}
