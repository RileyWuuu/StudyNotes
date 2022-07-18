package main

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

func TestSortSlice(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[]int{-4, -1, 0, 3, 10}},
			output{[]int{0, 1, 9, 16, 100}},
		},
		{
			input{[]int{-7, -3, 2, 3, 11}},
			output{[]int{4, 9, 9, 49, 121}},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, SortedSquare(x.i.in))
	}
}
