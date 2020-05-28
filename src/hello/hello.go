package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	leArquivo()
	introducao()

	for {
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
	sites := []string{"https://random-status-code.herokuapp.com/", "https://www.google.com", "https://www.alura.com.br"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site: ", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func testaSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site: ", site, "foi carregado com sucesso.")
	} else {
		fmt.Println("Site: ", site, "esta com problema. Status code: ", response.StatusCode)
	}
}

func leArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)

	linha, err := leitor.ReadString('\n')

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println(linha)

	return sites
}
