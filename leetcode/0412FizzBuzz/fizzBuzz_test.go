package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	in int
}

type output struct {
	out []string
}

type result struct {
	i input
	o output
}

func TestFizz(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{3},
			output{[]string{"1", "2", "Fizz"}},
		},
		{
			input{5},
			output{[]string{"1", "2", "Fizz", "4", "Buzz"}},
		},
		{
			input{15},
			output{[]string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
		},
	}

	for _, x := range data {
		ast.Equal(x.o.out, FizzBuzz(x.i.in))
	}
}
