package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	in []int
}
type output struct {
	out [][]int
}
type result struct {
	i input
	o output
}

func TestSortSlice(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[]int{-1, 0, 1, 2, -1, -4}},
			output{[][]int{{-1, -1, 2}, {-1, 0, 1}}},
		},
		{
			input{[]int{0, 1, 1}},
			output{[][]int{}},
		},
		{
			input{[]int{0, 0, 0}},
			output{[][]int{{0, 0, 0}}},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, threeSum(x.i.in))
	}
}
