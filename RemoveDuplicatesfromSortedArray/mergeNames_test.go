package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	in []string
}
type output struct {
	out []string
}
type result struct {
	i input
	o output
}

func TestMergeNames(t *testing.T) {
	ast := assert.New(t)
	fmt.Println("A")
	data := []result{
		{
			input{in: []string{"A", "B", "A", "A", "C", "Z", "H", "E", "B"}},
			output{out: []string{"A", "B", "C", "Z", "H", "E"}},
		},
		{
			input{in: []string{"T", "E", "S", "T", "I", "N", "G"}},
			output{out: []string{"T", "E", "S", "I", "N", "G"}},
		},
		{
			input{in: []string{"M", "E", "R", "G", "E", "N", "A", "M", "E"}},
			output{out: []string{"M", "E", "R", "G", "N", "A"}},
		},
	}
	for _, n := range data {
		ast.Equal(n.o.out, RemoveDuplicate(n.i.in))
	}
}
