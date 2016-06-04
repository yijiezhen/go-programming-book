package main

func main() {
	println(Join(",", "a", "b", "c", "d"))
	println(Join(",", "a"))
}

func Join(sep string, vals ...string) string {
	if (len(vals) == 0) {
		return ""
	}
	if (len(vals) == 1) {
		return vals[0]
	}
	n := len(sep) * (len(vals) - 1)
	for _, val := range vals {
		n += len(val)
	}

	data := make([]byte, n)

	bp := copy(data, vals[0])

	for i := 1; i < len(vals); i++ {
		bp += copy(data[bp:], sep)
		bp += copy(data[bp:], vals[i])
	}
	return string(data)
}
