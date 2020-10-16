package somemath

import (
	"errors"
)

//Fib - calculate fibonaci value
func Fib(n int) (int, error) {
	var a int
	var b int

	if n > 20 {
		return -1, errors.New("n parameter more then 20")
	}

	if n == 0 {
		return 0, nil
	}

	if n == 1 {
		return 1, nil
	}

	a, _ = Fib(n - 2)
	b, _ = Fib(n - 1)

	return a + b, nil
}
