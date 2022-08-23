package p0067

//TimeComplexity: O(n)
func addBinary(a string, b string) string {

	result := ""
	carry := 0
	x := len(a) - 1
	y := len(b) - 1

	for x >= 0 || y >= 0 {
		c := carry
		if x >= 0 && string(a[x]) == "1" {
			c += 1
		}
		if y >= 0 && string(b[y]) == "1" {
			c += 1
		}

		if c%2 == 1 {
			result = "1" + result
		} else {
			result = "0" + result
		}

		if c < 2 {
			carry = 0
		} else {
			carry = 1
		}
		x--
		y--
	}
	if carry != 0 {
		result = "1" + result
	}
	return result
}
