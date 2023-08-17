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

func TestSearch(t *testing.T) {
	ast := assert.New(t)

	data := []result{
		{
			input{in: []int{1, 2, 3, 1}},
			output{out: true},
		},
		{
			input{in: []int{1, 2, 3, 4}},
			output{out: false},
		},
		{
			input{in: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
			output{out: true},
		},
	}
	for _, n := range data {
		ast.Equal(n.o.out, n.i.in)
	}
}
