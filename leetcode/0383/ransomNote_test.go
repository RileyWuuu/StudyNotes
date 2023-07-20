package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	ransom   string
	magazine string
}
type output struct {
	out bool
}
type result struct {
	i input
	o output
}

func TestRansomNote(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{ransom: "a", magazine: "b"},
			output{false},
		},
		{
			input{ransom: "aa", magazine: "ab"},
			output{false},
		},
		{
			input{ransom: "aa", magazine: "aab"},
			output{true},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, RansomNote(x.i.ransom, x.i.magazine))
	}
}
