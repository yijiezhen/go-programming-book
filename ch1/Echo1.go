package main

import "fmt"
import "os"

func main() {
	var s string
	for i := 0 ; i < len(os.Args); i++ {
		s += fmt.Sprintf("%d: %s\n", i, os.Args[i])
	}
	fmt.Println(s)
}


