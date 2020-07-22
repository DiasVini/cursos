package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // Operação bloqueante
	fmt.Println("Só depois que foi lido")
}

func main() {
	c := make(chan int) // canal sem buffer

	go rotina(c)

	fmt.Println(<-c) // operação bloqueante - Se parece com um then do js
	fmt.Println("Foi lido")

	// Deadlock
	// Deadlock é ocasionado por tentarmos manipular um canal que já não está sendo usado
	//
	fmt.Println(<-c)

	fmt.Println("Fim")

}
