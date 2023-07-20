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

func TestDuplicate(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[]int{1, 0, 2, 3, 0, 4, 5, 0}},
			output{[]int{1, 0, 0, 2, 3, 0, 0, 4}},
		},
		{
			input{[]int{1, 2, 3}},
			output{[]int{1, 2, 3}},
		},
	}
	for _, x := range data {
		duplicateZeros(x.i.in)
		ast.Equal(x.o.out, x.i.in)
	}
}
