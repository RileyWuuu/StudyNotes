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

func TestRemove(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[]int{1, 1, 2}},
			output{2},
		},
		{
			input{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			output{5},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, removeDuplicates(x.i.in))
	}

}
