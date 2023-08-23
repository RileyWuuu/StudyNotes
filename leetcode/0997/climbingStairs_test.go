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

func TestClimb(t *testing.T) {
	ast := assert.New(t)

	data := []result{
		{
			input{in: 1},
			output{out: 1},
		},
		{
			input{in: 2},
			output{out: 2},
		},
		{
			input{in: 3},
			output{out: 3},
		},
	}
	for _, n := range data {
		ast.Equal(n.o.out, climbStairs(n.i.in))
	}

}
