package main

import "fmt"

func notaParaConceito(n float64) string {

	// Switch é para condições que não sejam só verdadeiro ou falso
	var nota = int(n)
	switch nota {
	case 10:
		// Palavra reservada para executar o proximo case
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota Invalida"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
	fmt.Println(notaParaConceito(11.1))
	fmt.Println(notaParaConceito(-2.1))
}
