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

func TestFinxMax(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[]int{1, 1, 0, 1, 1, 1}},
			output{3},
		},
		{
			input{[]int{1, 0, 1, 1, 0, 1}},
			output{2},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, FindMaxConsecutiveOnes(x.i.in))
		ast.Equal(x.o.out, FindMaxForRange(x.i.in))
	}
}
