package main

import (
	"fmt"
	"crypto/sha256"
)

func main() {
	s1 := "unodostresquatro"
	s2 := "UNODOSTRESQUATRO"
	h1 := sha256.Sum256([]byte(s1))
	h2 := sha256.Sum256([]byte(s2))
	fmt.Println(BitDiff(&h1, &h2))
}

func BitDiff(a, b *[sha256.Size]byte) (int) {
	n := 0;
	for i := range a {
		for c := a[i] ^ b[i]; c != 0; c &= c - 1 {
			n++
		}
	}
	return n
}