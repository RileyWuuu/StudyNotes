package p0118

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	in int
}
type output struct {
	out [][]int
}
type result struct {
	i input
	o output
}

func TestGenerate(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{5},
			output{[][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		},
		{
			input{1},
			output{[][]int{{1}}},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, generate(x.i.in))
	}
}
