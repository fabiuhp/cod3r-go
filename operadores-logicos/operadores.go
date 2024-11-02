package main

func compras(trabalho1, trabalho2 bool) (bool, bool, bool) {
	comprarTv50 := trabalho1 && trabalho2
	comprarTv32 := trabalho1 != trabalho2 // ou exclusivo
	comprarSorvete := trabalho1 || trabalho2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	println(tv50, tv32, sorvete)

	tv50, tv32, sorvete = compras(true, false)
	println(tv50, tv32, sorvete)

	tv50, tv32, sorvete = compras(false, true)
	println(tv50, tv32, sorvete)

	tv50, tv32, sorvete = compras(false, false)
	println(tv50, tv32, sorvete)
}
