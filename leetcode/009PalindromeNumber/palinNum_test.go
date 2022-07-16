package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	in int
}
type output struct {
	out bool
}
type result struct {
	i input
	o output
}

func TestPalindrome(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{in: 1234567},
			output{out: false},
		},
		{
			input{in: 78945654987},
			output{out: true},
		},
		{
			input{in: 99887766566778899},
			output{out: true},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, IsPalindrome(x.i.in))
	}

}
