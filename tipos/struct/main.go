package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco - p.preco*(p.desconto/100)
}

func main() {
	produto1 := produto{
		nome:     "Lápis",
		preco:    1.79,
		desconto: 5.0,
	}

	produto2 := produto{"Notebook", 1989.90, 10.0}

	fmt.Println(produto1, produto1.precoComDesconto())
	fmt.Println(produto2, produto2.precoComDesconto())
}
