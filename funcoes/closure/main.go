package main

func fatorial(n uint) uint {
	if n == 0 {
		return 1
	}

	return n * fatorial(n-1)
}

func main() {
	resultado := fatorial(5)
	println(resultado)
}
