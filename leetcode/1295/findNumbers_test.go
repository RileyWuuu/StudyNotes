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

func TestFindNum(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{[]int{12, 345, 2, 6, 7896}},
			output{2},
		},
		{
			input{[]int{555, 901, 482, 1771}},
			output{1},
		},
	}

	for _, x := range data {
		ast.Equal(x.o.out, FindNumbers(x.i.in))
		ast.Equal(x.o.out, FindNumbersSimple(x.i.in))
	}
}
