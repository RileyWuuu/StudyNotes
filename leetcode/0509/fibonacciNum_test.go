package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	n int
}

type output struct {
	result int
}

type testdata struct {
	in  input
	out output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{n: 1},
			output{result: 1},
		},
		{
			input{n: 2},
			output{result: 1},
		},
		{
			input{n: 3},
			output{result: 2},
		},
		{
			input{n: 4},
			output{result: 3},
		},
	}

	for _, data := range tds {
		ast.Equal(data.out.result, FibonacciNum(data.in.n))
	}
}
