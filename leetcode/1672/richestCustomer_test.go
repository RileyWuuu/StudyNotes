package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	in [][]int
}

type output struct {
	out int
}

type result struct {
	i input
	o output
}

func TestMaximumNum(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[][]int{{1, 2, 3}, {3, 2, 1}}},
			output{6},
		},
		{
			input{[][]int{{1, 5}, {7, 3}, {3, 5}}},
			output{10},
		},
		{
			input{[][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}},
			output{17},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, MaximumWealth(x.i.in))
	}
}
