package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(5, 97)
	if total != 102 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 102)
	}
}
