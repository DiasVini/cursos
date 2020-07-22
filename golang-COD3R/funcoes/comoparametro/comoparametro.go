package main

func multiplicacao(a, b int) int {
	return a * b
}

// devo declarar qual o tipo de função parametro e retorno esperado
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado := exec(multiplicacao, 3, 4)
	println(resultado)
}
