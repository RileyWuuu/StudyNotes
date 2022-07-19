package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums1 []int
	nums2 []int
	x     int
	y     int
}
type output struct {
	out []int
}
type result struct {
	i input
	o output
}

func TestMerge(t *testing.T) {
	ast := assert.New(t)

	data := []result{
		{
			input{
				nums1: []int{1, 2, 3, 0, 0, 0},
				nums2: []int{2, 5, 6},
				x:     3,
				y:     3,
			},
			output{[]int{1, 2, 2, 3, 5, 6}},
		},
		{
			input{
				nums1: []int{1},
				nums2: []int{},
				x:     1,
				y:     0,
			},
			output{[]int{1}},
		},
		{
			input{
				nums1: []int{0},
				nums2: []int{1},
				x:     0,
				y:     1,
			},
			output{[]int{1}},
		},
	}
	for _, x := range data {
		mergeII(x.i.nums1, x.i.x, x.i.nums2, x.i.y)
		ast.Equal(x.o.out, x.i.nums1)
	}
}
