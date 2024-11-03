package main

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // Campos anônimos
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	println(f.nome, f.velocidadeAtual, f.turboLigado)
}
