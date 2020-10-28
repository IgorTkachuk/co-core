package fib

import "testing"

func TestFib(t *testing.T) {

	if v := Calc(20); v != 6765 {
		t.Error("Expected 6765, got ", v)
	}
}
