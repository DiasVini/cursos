package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}

	// Blocos são obrigatorios
	// Parenteses não são necessarios
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.1)
}
