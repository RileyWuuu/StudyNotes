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

func TestFindDisappeared(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[]int{4, 3, 2, 7, 8, 2, 3, 1}},
			output{[]int{5, 6}},
		},
		{
			input{[]int{1, 1}},
			output{[]int{2}},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, findDisappearedNumbers(x.i.in))
	}
}
