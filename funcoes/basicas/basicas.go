package main

import "fmt"

func funcao1() {
	fmt.Println("Funcao 1")
}

func funcao2(p1 string, p2 string) {
	fmt.Printf("Funcao 2: %s %s\n", p1, p2)
}

func funcao3() string {
	return "Funcao 3"
}

func funcao4(p1, p2 string) string {
	return fmt.Sprintf("Funcao 4: %s %s", p1, p2)
}

func funcao5(p1 string, p2 string) (string, string) {
	return fmt.Sprintf("Funcao 5: %s", p1), fmt.Sprintf("Funcao 5: %s", p2)
}

func main() {
	funcao1()
	funcao2("parametro 1", "parametro 2")

	r3, r4 := funcao3(), funcao4("parametro 1", "parametro 2")
	fmt.Println(r3)
	fmt.Println(r4)

	r51, r52 := funcao5("parametro 1", "parametro 2")
	fmt.Println(r51)
	fmt.Println(r52)
}
