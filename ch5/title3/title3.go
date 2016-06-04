package main

import "fmt"

func main() {
	fmt.Printf("%s", title3())
}


func title3()  {
	type bailout struct{}
	defer func()  {
		switch p := recover(); p {
		case nil:
			panic(bailout{})
		default:
			panic(p)
		}
	}()
}