package main

import "fmt"

// Retorno nomeado
// A função que tem retorno nomeado pode usar o retorno  limpo
func trocar(p1, p2 int) (segundo int, primeiro int) {

	segundo = p2
	primeiro = p1
	return // Retorno limpo
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)

	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)

}
