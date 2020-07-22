// if com init não é nada mais que inicializar uma variavel e ou condição dentro da estrutura  de decisão do if
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {

	// rand - pacote de numeros aleatorios
	// NewSource passa uma fonte para um numero aleatorio
	// time.Now().UnixNano() pega o milisegundo atual
	s := rand.NewSource(time.Now().UnixNano())
	// Gerar um numero novo
	r := rand.New(s)
	// Intn gera um numero aleatorio ate 10
	return r.Intn(10)
}

func main() {
	// if init
	// Tambem suportado pelo switch
	// Usou a ideia do for e colocou no if e no switch
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("Ganhou!!!")
	} else {
		fmt.Println("Perdeu!")
	}
}
