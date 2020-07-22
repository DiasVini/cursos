package main

import "fmt"

func main() {
	// Outra forma de inicializar um map
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15464.78,
		"Pedro Junior":   1200.0,
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Inexistente") //Não gera erro

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
