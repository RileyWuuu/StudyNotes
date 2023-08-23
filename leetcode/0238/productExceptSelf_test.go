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

func TestProductExceptSelf(t *testing.T) {
	ast := assert.New(t)

	data := []result{
		{
			input{in: []int{1, 2, 3, 4}},
			output{out: []int{24, 12, 8, 6}},
		},
		{
			input{in: []int{-1, 1, 0, -3, 3}},
			output{out: []int{0, 0, 9, 0, 0}},
		},
	}
	for _, n := range data {
		ast.Equal(n.o.out, productExceptSelf(n.i.in))
	}
}
