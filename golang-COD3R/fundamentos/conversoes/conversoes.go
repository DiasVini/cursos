package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))
	fmt.Println(int(x) / y)

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado..
	// string(97) seria igual a letra a, pois os valores seriam buscados na tabela ascii
	fmt.Println("Teste " + string(97))

	// int  para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// String para int
	// Essa função retorna dois valores, por isso duas variaveis
	// Primeiro valor é nosso resultado, segunda é um erro
	// Para ignorar a variavel é so usar _
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 123)

	// podemos usar true, false, 0, 1
	b, _ := strconv.ParseBool("true")

	if b {
		fmt.Println("verdadeiro")
	} else {
		fmt.Println("Falso")
	}

}
