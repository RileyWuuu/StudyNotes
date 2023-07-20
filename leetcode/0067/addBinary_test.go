package p0067

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	a string
	b string
}
type output struct {
	out string
}
type result struct {
	i input
	o output
}

func TestAddBinary(t *testing.T) {
	ast := assert.New(t)

	data := []result{
		{
			input{a: "11", b: "1"},
			output{"100"},
		},
		{
			input{a: "1010", b: "1011"},
			output{"10101"},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, addBinary(x.i.a, x.i.b))
	}
}
