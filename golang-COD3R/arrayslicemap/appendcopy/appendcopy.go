package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	// Pegar tudo que tem no slice1 e copiar para o slice2
	// como o tamanho do slice1 é maior que o slice2 o copy pegou os primeiros elementos do slice1
	copy(slice2, slice1)
	fmt.Println(slice2)
}
