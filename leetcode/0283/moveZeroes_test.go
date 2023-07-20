package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	in []int
}
type output struct {
	output []int
}
type result struct {
	i input
	o output
}

func TestMoveZeroes(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[]int{0, 1, 0, 3, 12}},
			output{[]int{1, 3, 12, 0, 0}},
		},
		{
			input{[]int{0}},
			output{[]int{0}},
		},
		{
			input{[]int{4, 5, 0, 1, 0, 9, 6}},
			output{[]int{4, 5, 1, 9, 6, 0, 0}},
		},
	}
	for _, x := range data {
		moveZeroes(x.i.in)
		ast.Equal(x.o.output, x.i.in)
	}

}
