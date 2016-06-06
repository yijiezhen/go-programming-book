package main

import (
	"bytes"
	"io"
)



func main() {
	var buf *bytes.Buffer
	println(buf == nil)
	var buf2 io.Writer
	println(buf2 == nil)
}


