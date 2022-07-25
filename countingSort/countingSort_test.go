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

func TestCount(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[]int{9, 4, 1, 5, 6, 2, 2, 7, 1}},
			output{[]int{1, 1, 2, 2, 4, 5, 6, 7, 9}},
		},
		{
			input{[]int{8, 4, 6, 7, 1, 9, 3, 4, 2, 7, 1}},
			output{[]int{1, 1, 2, 3, 4, 4, 6, 7, 7, 8, 9}},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, countSort(x.i.in))
	}
}
