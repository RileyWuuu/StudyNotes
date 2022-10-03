package p2427

func commonFactors(a int, b int) int {
	if a == 1 && b == 1 {
		return 1
	}
	result := 0
	for i := 1; i <= a && i <= b; i++ {
		if a%i == 0 && b%i == 0 {
			result += 1
		}
	}
	return result
}
