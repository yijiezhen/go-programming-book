package main

import (
	"fmt"
	"bytes"
)

func main() {
	fmt.Println(comma1("1"))
	fmt.Println(comma1("123"))
	fmt.Println(comma1("1234"))
	fmt.Println(comma1("12345"))
	fmt.Println(comma1("123456"))
	fmt.Println(comma1("1234567"))

	fmt.Println(comma2("1"))
	fmt.Println(comma2("123"))
	fmt.Println(comma2("1234"))
	fmt.Println(comma2("12345"))
	fmt.Println(comma2("123456"))
	fmt.Println(comma2("1234567"))
}

func comma1(s string) string {
	length := len(s)
	if length <= 3 {
		return s
	}
	return comma1(s[:length - 3]) + "," + s[length - 3:]
}

func comma2(s string) string {
	length := len(s)
	if length <= 3 {
		return s
	}
	var buf bytes.Buffer
	pre := length % 3
	buf.WriteString(s[:pre])
	for i := pre; i + 3 <= length; i += 3 {
		if i > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(s[i:i+3])
	}
	return buf.String()
}
