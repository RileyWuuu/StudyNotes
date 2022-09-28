package p1680

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

func TestconcatenatedBinary(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{1},
			output{1},
		},
		{
			input{3},
			output{27},
		},
		{
			input{12},
			output{505379714},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, concatenatedBinary(x.i.in))
	}
}
