package main

import (
	"fib/somemath"
	"flag"
	"fmt"
)

func main() {

	posPtr := flag.Int("n", 0, "fib position")
	helpPtr := flag.Bool("h", false, "get help")
	flag.Parse()

	if *helpPtr {
		fmt.Println("Usage main.exe [-n <number>] [-h]")
	} else {
		res, err := somemath.Fib(*posPtr)

		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}

		fmt.Println(res)
	}
}
