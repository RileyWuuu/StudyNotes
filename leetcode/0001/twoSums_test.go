package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	in     []int
	target int
}
type output struct {
	out []int
}
type result struct {
	i input
	o output
}

func TestTwoSums(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[]int{2, 7, 11, 15}, 9},
			output{[]int{0, 1}},
		},
		{
			input{[]int{3, 2, 4}, 6},
			output{[]int{1, 2}},
		},
		{
			input{[]int{3, 3}, 6},
			output{[]int{0, 1}},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, TwoSums(x.i.in, x.i.target))
	}
}
