package fib

//Calc - calculate fibonaci value
func Calc(n int) int {

	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}

	return Calc(n-2) + Calc(n-1)
}
