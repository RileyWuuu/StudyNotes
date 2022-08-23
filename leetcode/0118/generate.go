package p0118

//TimeComplextiy:O(n^2)
func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		result[i][i] = 1
		for x := 1; x < i; x++ {
			result[i][x] = result[i-1][x-1] + result[i-1][x]
		}
	}
	return result
}
