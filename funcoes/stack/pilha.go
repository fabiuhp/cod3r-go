package main

import "runtime/debug"

func funcao3() {
	debug.PrintStack()
}

func funcao2() {
	funcao3()
}

func funcao1() {
	funcao2()
}

func main() {
	funcao1()
}
