package main

import "io"

type LimitReader struct {
	R io.Reader
	N int64
}

func (l *LimitReader) Read(p []byte) (n int, err error) {
	if l.N <=0 {
		return 0, io.EOF
	}
	if l.N < len(p) {
		p = p[0:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}

func main() {

}
