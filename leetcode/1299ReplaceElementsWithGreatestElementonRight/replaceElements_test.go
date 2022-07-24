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

func TestReplace(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[]int{17, 18, 5, 4, 6, 1}},
			output{[]int{18, 6, 6, 6, 1, -1}},
		},
		{
			input{[]int{400}},
			output{[]int{-1}},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, replaceElements(x.i.in))
	}

}
