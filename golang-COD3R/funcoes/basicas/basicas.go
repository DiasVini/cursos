package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

// Parametros do mesmo tipo
func f4(p1, p2 string) string {
	// Sprintf retorna uma string formatada
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

// Multiplos retornos
func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("Vini", "Dias")
	fmt.Println(f3())
	fmt.Println(f4("Dias", "Vini"))

	r51, r52 := f5()

	fmt.Println("F5: ", r51, r52)
}
