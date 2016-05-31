package main

import "strconv"

type A struct {
	a int
}

func main() {
	var s = []string{"1", "2"}
	println(s)
	testB(s)
	println(s)
}

func testA(a *A) {
	a.a = 20
}

func testB(stack []string) {
	println(stack)
	stack = append(stack, strconv.Itoa(len(stack) + 1))
	println(stack)
}
