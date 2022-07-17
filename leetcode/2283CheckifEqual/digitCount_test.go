package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	in string
}
type output struct {
	out bool
}
type result struct {
	i input
	o output
}

func TestDigitCount(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{"1210"},
			output{true},
		},
		{
			input{"030"},
			output{false},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, DigitCount(x.i.in))
	}
}
