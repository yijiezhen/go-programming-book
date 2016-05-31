package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "need two ints for gcd calculation")
		os.Exit(1)
	}
	a, err := strconv.Atoi(os.Args[1])
	 if err != nil {
		fmt.Printf("can't parse %v to an integer", os.Args[1])
		os.Exit(1)
	}

	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("can't parse %v to an integer", os.Args[2])
		os.Exit(1)
	}

	g := gcd(a, b)

	fmt.Printf("gcd(%d,%d)=%d\n", a, b, g)
}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b;
	}
	return a;
}
