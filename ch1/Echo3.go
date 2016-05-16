package main

import (
	"fmt"
	"strings"
)

func main() {

}

func slow(n int) {
	var s string
	for i := 0; i < n; i++ {
		s += fmt.Sprintf("%d \n", i)
	}
	fmt.Printf(s)
}
