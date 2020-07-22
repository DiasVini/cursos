package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	// Switch sem condição vai procurar o primeiro case que tenha uma expressão que seja verdadeira
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa Tarde.")
	default:
		fmt.Println("Boa Noite.")
	}
}
