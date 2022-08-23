package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	in int
}
type output struct {
	out int
}
type result struct {
	i input
	o output
}

func TestNumsOfSteps(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{14},
			output{6},
		},
		{
			input{8},
			output{4},
		},
		{
			input{123},
			output{12},
		},
	}

	for _, x := range data {
		ast.Equal(x.o.out, NumberOfSteps(x.i.in))
	}
}
