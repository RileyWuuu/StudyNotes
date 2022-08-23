package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	in []int
}

type output struct {
	out bool
}

type result struct {
	i input
	o output
}

func TestCheckIfExist(t *testing.T) {
	ast := assert.New(t)

	data := []result{
		{
			input{[]int{10, 2, 5, 3}},
			output{true},
		},
		{
			input{[]int{7, 1, 14, 11}},
			output{true},
		},
		{
			input{[]int{3, 1, 7, 11}},
			output{false},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, checkIfExist(x.i.in))
		ast.Equal(x.o.out, checkIfExistII(x.i.in))
	}
}
