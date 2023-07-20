package p0066

func plusOne(digits []int) []int {
	len := len(digits)
	for i := len-1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	twoDigits := make([]int, len+1)
	twoDigits[0] = 1
	return twoDigits
}
