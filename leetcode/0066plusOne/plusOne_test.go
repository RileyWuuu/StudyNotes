package p0066

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

func TestPlusOne(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[]int{1, 2, 3}},
			output{[]int{1, 2, 4}},
		},
		{
			input{[]int{4, 3, 2, 1}},
			output{[]int{4, 3, 2, 2}},
		},
		{
			input{[]int{9}},
			output{[]int{1, 0}},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, plusOne(x.i.in))
	}
}
