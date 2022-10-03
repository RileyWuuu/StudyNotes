package p2427

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	a int
	b int
}
type output struct {
	out int
}

type result struct {
	i input
	o output
}

func TestCommonFactors(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{a: 12, b: 6},
			output{4},
		},
		{
			input{a: 25, b: 30},
			output{2},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, commonFactors(x.i.a, x.i.b))
	}
}
