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

func TestThirdMax(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[]int{3, 2, 1}},
			output{1},
		},
		{
			input{[]int{1, 2}},
			output{2},
		},
		{
			input{[]int{2, 2, 3, 1}},
			output{1},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, thirdMax(x.i.in))
	}

}
