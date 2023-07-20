package p0948

import "sort"

func bagOfTokensScore(tokens []int, power int) int {
	low, high := 0, len(tokens)-1
	sort.Ints(tokens)
	point := 0
	for low <= high {
		if power-tokens[low] >= 0 {
			point++
			power -= tokens[low]
			low++
		} else if point > 0 && low < high {
			point--
			power += tokens[high]
			high--
		} else {
			break
		}
	}
	return point
}
