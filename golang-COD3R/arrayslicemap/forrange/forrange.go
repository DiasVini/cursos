package main

import "fmt"

func main() {
	// iniciar um array do tipo int, e o numero de indices é definido pela quantidade de valores que eu adicionar na inicialização
	numeros := [...]int{1, 2, 3, 4, 5} // O compilador conta!

	// Esse for retorna dois valores
	// i é o indice atual
	// numero é o valor que esta sendo percorrido no array
	// range pega a quantidade de indices relacionados a um vetor/array/slice
	// igual ao ForEach do JS
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	// ignorando o indice
	for _, num := range numeros {
		fmt.Println(num)
	}

	// se eu suprimir um dos valores, o retornado é o indice
	for i := range numeros {
		fmt.Println(i)
	}
}
