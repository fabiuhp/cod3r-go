package main

func multiplicacao(a, b int) int {
	return a * b
}

func execucao(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado := execucao(multiplicacao, 3, 4)
	println(resultado)
}
