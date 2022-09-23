package p1680

import "fmt"

func main() {
	fmt.Println(concatenatedBinary(5))
}

func concatenatedBinary(n int) int {
	ans := 1
	cur, count := 2, 1
	for i := 2; i <= n; i++ {
		if i == cur {
			cur <<= 1
			count++
		}
		ans <<= count
		ans += i
		ans %= 1000000007
	}
	return ans

}
