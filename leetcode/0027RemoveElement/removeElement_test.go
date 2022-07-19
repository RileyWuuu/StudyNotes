package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	in  []int
	val int
}
type output struct {
	out int
}
type result struct {
	i input
	o output
}

func TestRemoveElement(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{in: []int{3, 2, 2, 3}, val: 3},
			output{2},
		},
		{
			input{in: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2},
			output{5},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, removeElement(x.i.in, x.i.val))
	}
}
