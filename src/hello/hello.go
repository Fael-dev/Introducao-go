package main

import "fmt"

func main() {

	nome := "Rafael"
	versao := 1.1
	fmt.Println("Olá Sr.", nome)
	fmt.Println("A versão do programa é ", versao)
	fmt.Println("\n1 - Iniciar monitoramento.")
	fmt.Println("2 - Exibir logs.")
	fmt.Println("0 - Sair do programa.")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("O endereço da minha variável comando é ", &comando)
	fmt.Println("O comando escolhido foi ", comando)

	if comando == 1 {
		fmt.Println("\nMonitorando...")
	} else if comando == 2 {
		fmt.Println("\nExibindo logs...")
	} else if comando == 0 {
		fmt.Println("\nSaindo do programa.")
	} else {
		fmt.Println("\nNão conheço este comando.")
	}

}
