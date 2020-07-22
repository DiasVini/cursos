package main

import "fmt"

// função variatica recebe um numero indefinido de parametros
func media(numeros ...float64) float64 {
	total := 0.0

	for _, num := range numeros {
		total += num
	}

	return total
}

func main() {
	fmt.Printf("Média: %.2f", media(7.7, 8.1, 5.9))
	fmt.Printf("Média: %.2f", media(7.7, 8.1))
	fmt.Printf("Média: %.2f", media(7.7, 8.1, 5.9, 9.2, 14.5, 58.3))

}
