package main

import "fmt"

// quando passamos uma variavel para uma função atraves de parametros, criamos uma copia do valor
// a finalidade de se passar o ponteiro para função é justamente alterar a variavel em si
// Acessar a memoria na qual ela esta armazenada

func inc1(n int) {
	n++
}

// revisão: um ponteiro é representado por um *
func inc2(n *int) {
	// revisão: * é usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) // por valor
	fmt.Println(n)

	// revisão: & é usado para obter o endereço da variavel
	inc2(&n) // por referencia
	fmt.Println(n)

}
