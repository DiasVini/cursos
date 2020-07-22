package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Pq vc não fala comigo?", 3)
	// fale("João", "Só posso falar depois de vc!", 1)

	// Tipo uma thread porem mais leve e com menos recursos
	// uma thread pode gerenciar varias goroutines
	// Thread é uma linha de execução independente

	// A go routine só vai continuar executando se a thread principal estiver sendo executada

	// goroutine
	go fale("Maria", "Entendi", 10)
	// Linha de Execução principal
	fale("João", "Parabens", 5)
}
