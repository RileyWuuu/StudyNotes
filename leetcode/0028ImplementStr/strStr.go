package p0028

//Runtime: 3ms
func strStr(haystack string, needle string) int {

	x, y := len(haystack), len(needle)
	for i := 0; i <= x-y; i++ {
		if haystack[i:i+y] == needle {
			return i
		}
	}

	return -1
}
