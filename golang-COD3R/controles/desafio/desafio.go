package main

import "fmt"

func notaParaConceito(n float64) string {
	nota := int(n)

	switch {
	case nota == 10, nota == 9:
		return "A"
	case nota == 7, nota == 8:
		return "B"
	case nota == 5, nota == 6:
		return "C"
	case nota == 3, nota == 4:
		return "D"
	case nota == 2, nota == 1, nota == 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.8))
	fmt.Println(notaParaConceito(5.4))
	fmt.Println(notaParaConceito(3.8))
	fmt.Println(notaParaConceito(-1.8))
}
