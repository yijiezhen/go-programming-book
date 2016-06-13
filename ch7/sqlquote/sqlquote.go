package main

import "fmt"

func main() {

}


func sqlQuote(x interface{}) string {
	switch x.(type) {
	case nil: return "NULL"
	case int, uint: return fmt.Sprintf("%d", x)
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return x
	default:
		panic("")
	}
}
