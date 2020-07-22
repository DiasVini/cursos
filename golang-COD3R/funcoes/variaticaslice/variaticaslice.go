package main

import "fmt"

// Passando um slice como parametro variavel
// Eu trato esse parametro como array
func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilerme"}

	// Os tres pontinhos (...) dentro das chaves, conta as entradas e cria um array
	// aprovados := [...]string{"maria", "Pedro", "Rafael", "Guilerme"}

	// passar slice com ... apos o nome do slice
	imprimirAprovados(aprovados...)
}
