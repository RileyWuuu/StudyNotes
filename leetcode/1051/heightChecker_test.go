package main

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

func TestHeightCheck(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[]int{1, 1, 4, 2, 1, 3}},
			output{3},
		},
		{
			input{[]int{5, 1, 2, 3, 4}},
			output{5},
		},
		{
			input{[]int{1, 2, 3, 4, 5}},
			output{0},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, heightChecker(x.i.in))
		ast.Equal(x.o.out, heightCheckerII(x.i.in))
	}
}
