package main

import "strings"

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) nomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{"João", "Pedro"}
	println(p1.nomeCompleto())

	p1.setNomeCompleto("Maria Pereira")
	println(p1.nomeCompleto())
}
