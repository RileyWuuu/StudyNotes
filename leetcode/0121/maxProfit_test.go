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

func TestMaxProfit(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[]int{7,1,5,3,6,4}},
			output{5},
		},
		{
			input{[]int{7,6,4,3,1}},
			output{0},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, TwoSums(x.i.in, x.i.target))
	}
}
