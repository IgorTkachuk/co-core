package somemath

import "testing"

func TestFib(t *testing.T) {
	var v int

	v, _ = Fib(20)
	if v != 6765 {
		t.Error("Expected 6765, got ", v)
	}
}
