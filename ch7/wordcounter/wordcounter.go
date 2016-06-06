package main

import (
	"bufio"
	"fmt"
	"bytes"
)

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewReader(p))
	s.Split(bufio.ScanWords)
	n := 0
	for {
		r := s.Scan(); if s.Err() != nil {
			panic("error to read")
		}
		if r == true {
			n++
		} else {
			break
		}
	}
	*w += WordCounter(n)
	return n, nil
}



func main() {
	var w WordCounter
	w.Write([]byte("hello world"))
	fmt.Println(w)
}
