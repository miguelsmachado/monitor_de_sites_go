package main

import (
	"fmt"
)

func main() {
	var nome string = "Miguel"
	versao := 1.1

	fmt.Println("Olá sr.", nome)
	fmt.Println("Esse programa está na versão", versao)

	fmt.Println("\n1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir os logs")
	fmt.Println("0 - Sair do programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	// Usando switch (in case of do pascal)
	switch comando{
	case 1: 
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa")
	default:
		fmt.Println("Não conheço este comando")
	}
}
