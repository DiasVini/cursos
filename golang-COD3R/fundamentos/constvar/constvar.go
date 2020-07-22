package main

import (
	"fmt"
	// Referenciar importação com outro nome
	m "math"
	// Referenciar importação que não será usada
	_ "net"
)

func main() {
	// Forma padrão de declarar constante
	// float64 tipo padrão de um literal com ponto flutuante
	const PI float64 = 3.1415
	// Forma padrão da variavel
	var raio = 3.2
	// Forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	// Declarar variaveis e constantes em blocos
	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	// Declarar multiplas variaves em uma linha
	var e, f bool = true, false
	fmt.Println(e, f)

	// Declarar multiplas variaveis com operador de atribuição
	g, h, i := 2, false, "Epa!"
	fmt.Println(g, h, i)

	// Cada variavel em Go não muda seu tipo

}
