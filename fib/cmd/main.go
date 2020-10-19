package main

import (
	"flag"
	"fmt"

	"itka4yk.go/fib"
)

func main() {

	posPtr := flag.Int("n", 0, "fib position")
	flag.Parse()

	if *posPtr > 20 {
		fmt.Printf("n parameter more then 20")
		return
	}

	fmt.Println(fib.Calc(*posPtr))
}
