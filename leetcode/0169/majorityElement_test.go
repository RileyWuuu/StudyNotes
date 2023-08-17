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

func TestMajorityElement(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[]int{3, 2, 3}},
			output{3},
		},
		{
			input{[]int{2, 2, 1, 1, 1, 2, 2}},
			output{2},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, majorityElement(x.i.in))
	}
}
