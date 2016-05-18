package main

func main() {

}

func gcd(a int, b int) {
	for b != 0 {
		a, b = b, a%b;
	}
	return a;
}
