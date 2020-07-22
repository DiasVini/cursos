package main

import "fmt"

// Defer é uma palavra reservada, que executa um trecho de codigo pouco antes do return, ou do final da função
// Pode ser usado, por exemplo, para fechar uma conexão com o banco de dados
func obterValorAprovado(numero int) int {
	defer fmt.Println("fim!")
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero

}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
