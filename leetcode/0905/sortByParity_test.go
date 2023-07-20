package p0905

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	in []int
}
type output struct {
	out []int
}

type result struct {
	i input
	o output
}

func TestSortArrayByParity(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[]int{3, 1, 2, 4}},
			output{[]int{2, 4, 3, 1}},
		},
		{
			input{[]int{0}},
			output{[]int{0}},
		},
	}
	for _, x := range data {
		ast.Eqaual(x.o.out, SortArrayByParity(x.i.in))
	}
}
