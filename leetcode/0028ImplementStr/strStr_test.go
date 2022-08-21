package p0028

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	haystack string
	needle   string
}
type output struct {
	out int
}
type result struct {
	i input
	o output
}

func TestStr(t *testing.T) {
	ast := assert.New(t)
	data := []result{
		{
			input{haystack: "hello", needle: "ll"},
			output{2},
		},
		{
			input{haystack: "aaaaa", needle: "bba"},
			output{-1},
		},
	}
	for _, x := range data {
		ast.Equal(x.o.out, strStr(x.i.haystack, x.i.needle))
	}
}
