package main

import "fmt"

// Compras é uma operação binaria pois tem dois parametros
func compras(trab1, trab2 bool) (bool, bool, bool) {

	comprarTv50 := trab1 && trab2
	// não existe operador "ou" exclusivo em Go
	comprarTv32 := trab1 != trab2 // Equivalente ao ou exclusivo
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudavel: %t", tv50, tv32, sorvete, !sorvete)
}
