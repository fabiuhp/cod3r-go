package main

var soma = func(a, b int) int {
	return a + b
}

func main() {
	println(soma(2, 3))

	sub := func(a, b int) int {
		return a - b
	}

	println(sub(2, 3))
}
