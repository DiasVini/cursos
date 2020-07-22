package main

import "fmt"

func main() {
	// Slice com 10 posições sem referencia a nenhum array
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	//  Cria um slice com 10 posições e cria um array com 20 posições que sera referenciado
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	// Adiciona posições a um slice
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	// o Slice referencia mais arrays quando passamos o limite do array referenciado
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))

}
