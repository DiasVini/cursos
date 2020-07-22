package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan - canal somente-leitura

func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)             // Pega a pagina
			html, _ := ioutil.ReadAll(resp.Body) // usa o pacote ioutil para manipulação de texto

			r, _ := regexp.Compile("<title>(.*?)<\\/title>") // aplica um regex ao texto lido pelo ReadAll
			c <- r.FindStringSubmatch(string(html))[1]       // Aplica o filtro de texto e retorna o valor encontrado

		}(url) // modo de passar parametro para uma Função anonima
	}
	return c
}

func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.agenciaspasso.com.br", "https://www.youtube.com")

	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
