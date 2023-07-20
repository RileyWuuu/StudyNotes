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

func TestSumArray(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[]int{1, 2, 3, 4}},
			output{[]int{1, 3, 6, 10}},
		},
		{
			input{[]int{1, 1, 1, 1, 1}},
			output{[]int{1, 2, 3, 4, 5}},
		},
		{
			input{[]int{3, 1, 2, 10, 1}},
			output{[]int{3, 4, 6, 16, 17}},
		},
	}

	for _, x := range data {
		ast.Equal(x.o.out, RunningSum(x.i.in))
	}

}
