package main

import "fmt"

func main() {

	// Canal é usado para comunicação entre goroutines
	// Channel é um tipo está dentro da propria estrutura da linguagem (int, e etc)
	// podemos trafegar dados dentro da channel

	// Criar um canal de comunicação que passa inteiros
	// o segundo paramentro é um buffer
	ch := make(chan int, 1)

	// Enviando o valor 1 pelo canal que aceita um valor inteiro (escrita)
	ch <- 1

	// recebendo dados do canal (leitura)
	<-ch

	ch <- 2
	fmt.Println(<-ch)

	// Canal é um instrumento de sincronismo
}
