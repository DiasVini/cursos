package main

import (
	"fmt"
	"time"
)

// Go é fortemente tipado
// interface é um tipo generico para pasar como parametro
func tipo(i interface{}) string {

	// i.(type) pega o tipo de parametro que foi passado
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "Função"
	default:
		return "Não sei"
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo("Olá"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
