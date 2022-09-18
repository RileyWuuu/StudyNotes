package p0948

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	token []int
	power int
}
type output struct {
	out int
}

type result struct {
	i input
	o output
}

func TestbagOfTokens(t *testing.T) {
	asc := assert.New(t)

	data := []result{
		{
			input{token: []int{100}, power: 50},
			output{0},
		},
		{
			input{token: []int{100, 200}, power: 150},
			output{1},
		},
		{
			input{token: []int{100, 200, 300, 400}, power: 200},
			output{2},
		},
	}
	for _, x := range data {
		asc.Equal(x.o.out, bagOfTokensScore(x.i.token, x.i.power))
	}
}
