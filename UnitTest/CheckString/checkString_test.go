package main

import "testing"

func TestString(t *testing.T) {
	expected := "Hello Nancy!"
	actual := Hello("Nancy")
	if actual != expected{
		t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
	}
}
