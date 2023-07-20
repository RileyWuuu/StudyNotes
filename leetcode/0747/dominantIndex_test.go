package p0747

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

func TestDominant(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[]int{3, 6, 1, 0}},
			output{1},
		},
		{
			input{[]int{1, 2, 3, 4}},
			output{-1},
		},
		{
			input{[]int{2, 3, 4, 9, 9}},
			output{-1},
		},
	}

	for _, x := range data {
		ast.Equal(x.o.out, dominantIndex(x.i.in))
	}
}
