package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Quando eu defino o tamanho de um vetor eu tenho um array
	// Quando eu não defino o tamanho de um vetor eu tenho um slice - Vetor de tamanho dinamico
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	// O slice é uma estrutura com um tamanho e aponta para outro elemento
	// posso ter um array interno
	// e varios slices apontando pra ele, com parte desse array
	// Slice não é um array! Slice define um pedaço de um array.
	// Slice é uma estrutura continua
	// É um estrutra com um ponteiro para o primeiro elemento e um tamanho
	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3]

	fmt.Println(a2, s2)

	// vai do indice 0 ao 2
	s3 := a2[:2]

	fmt.Println(a2, s3)

	// vc pode imaginar um slice como: estrutura com Tamanho e um ponteiro para um elemento de um array
	s4 := s2[:1]

	fmt.Println(s2, s4)

}
