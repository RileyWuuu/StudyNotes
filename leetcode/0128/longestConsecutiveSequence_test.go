package p0128

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	in []int
}
type output struct {
	out int
}
type result struct {
	i input
	o output
}

func TestLongestConsecutive(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[]int{100, 4, 200, 1, 3, 2}},
			output{3},
		},
		{
			input{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}},
			output{9},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, longestConsecutive(x.i.in))
	}
}
