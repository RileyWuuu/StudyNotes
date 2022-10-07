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
	out int
}
type result struct {
	i input
	o output
}

func TestSearch(t *testing.T) {
	ast := assert.New(t)

	data := []result{
		{
			input{in: []int{1, 3, 5, 6}, target: 5},
			output{out: 2},
		},
		{
			input{in: []int{1, 3, 5, 6}, target: 2},
			output{out: 1},
		},
		{
			input{in: []int{1, 3, 5, 6}, target: 7},
			output{out: 4},
		},
	}
	for _, n := range data {
		ast.Equal(n.o.out, SearchInsert(n.i.in, n.i.target))
	}
}
