package main

import "fmt"

// Atribuir função a variavel
var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))

	// Posso criar funções anonimas dentro de outra função
	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2, 3))
}
