package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	introducao()
	menu()

	comando := leComando()
	switch comando {
	case 1:
		monitoramento()
	case 2:
		fmt.Println("\nExibindo logs...")
	case 0:
		fmt.Println("\nSaindo do programa.")
		os.Exit(0)
	default:
		fmt.Println("\nNão conheço este comando.")
		os.Exit(-1)
	}
}

func introducao() {
	nome := "Rafael"
	versao := 1.1
	fmt.Println("Olá Sr.", nome)
	fmt.Println("A versão do programa é ", versao)
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi ", comando)

	return comando
}

func menu() {
	fmt.Println("\n1 - Iniciar monitoramento.")
	fmt.Println("2 - Exibir logs.")
	fmt.Println("0 - Sair do programa.")
}

func monitoramento() {
	fmt.Println("\nMonitorando...")
	site := "https://www.alura.com.br"
	response, _ := http.Get(site)
	fmt.Println(response)
}
