package main

import "fmt"

func main() {

	// maps é tipo um array com indices, com chave e valor
	// var aprovados map[int]string
	// Mapas devem ser inicializados
	aprovados := make(map[int]string)
	aprovados[12345678978] = "Maria"
	aprovados[12334523456] = "Pedro"
	aprovados[98765432112] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678978])
	// Apagar entrada de um map
	delete(aprovados, 12345678978)
	fmt.Println(aprovados[12345678978])
}
