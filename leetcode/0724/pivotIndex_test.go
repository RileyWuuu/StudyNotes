package p0724

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

func TestPivot(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[]int{1, 7, 3, 6, 5, 6}},
			output{3},
		},
		{
			input{[]int{1, 2, 3}},
			output{-1},
		},
		{
			input{[]int{2, 1, -1}},
			output{0},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, pivotIndex(x.i.in))
	}

}
