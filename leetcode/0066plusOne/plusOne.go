package p0066

func plusOne(digits []int) []int {
	len := len(digits) - 1
	for i := len; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	a := make([]int, len+2)
	a[0] = 1
	return a
}
