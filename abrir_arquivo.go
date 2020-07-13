package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"bufio"
	"io"
	"strings"
)

func main(){

	fmt.Println("")
	leArquivo0()
	fmt.Println("")
	leArquivo1()
	fmt.Println("")
	leArquivo2()
	fmt.Println("")
	leArquivo3()
	fmt.Println("")
	fmt.Println(leArquivo4())
}

func leArquivo0() [] string {

	var sites []string

	// A bibilioteco os, após ler o arquivo, devolve apenas o endereço de 
	// memória em que o arquivo se encontra
	arquivo, err  := os.Open("sites.txt")

	fmt.Println(arquivo)
	fmt.Println(err)

	return sites
}

func leArquivo1() [] string {

	var sites []string

	// A biblioteca io/ioutil, é capaz de ler o arquivo todo de uma só vez
	// É necessário realizar a conversão para string
	arquivo, err  := ioutil.ReadFile("sites.txt")

	fmt.Println(string(arquivo))
	fmt.Println(err)

	return sites
}

func leArquivo2() [] string {

	var sites []string
	// Nesse processo, vamos ler apenas uma linha

	// Abrimos o arquivo com o os
	arquivo, err  := os.Open("sites.txt")

	// Usando a biblioteca bufio e o método NewReader, vamos gerar um leitor
	// passamos o arquivo como parametro
	leitor := bufio.NewReader(arquivo)

	// Um objeto do tipo NewReader possui um método chamado ReadString
	// É preciso informar o byte limitador. Para infomrar o byte, passamos entre
	// aspas simples ''
	linha, err := leitor.ReadString('\n')

	// Tratamento de erro
	if err != nil{
		fmt.Println("Ocorreu um erro", err)
	}
	fmt.Println(linha)

	return sites
}

func leArquivo3() [] string {

	var sites []string
	// Nesse processo, vamos ler linha por linha

	// Abrimos o arquivo com o os
	arquivo, _ := os.Open("sites.txt")

	// Usando a biblioteca bufio e o método NewReader, vamos gerar um leitor
	// passamos o arquivo como parametro
	leitor := bufio.NewReader(arquivo)

	// Aqui, vamos usar a mesma técnica da função acima, mas desse vez fazendo um laço
	// for para ler linha por linha

	for {
		linha, err := leitor.ReadString('\n')
		// Com a biblioteca strings, podemos eliminar os espaços do final de cada linha
		// Como o strip no python

		linha = strings.TrimSpace(linha)
		fmt.Println(linha)
		// vamos usar o erro EOF (End of File) para encerrar o loop quando não houver mais nada para ser lido
		if err == io.EOF{
			break
		} else if err != nil{
			fmt.Println("Ocorreu um erro", err)
		}
		
	}
	
	return sites
}

func leArquivo4() [] string {

	var sites []string
	arquivo, _ := os.Open("sites.txt")
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF{
			break
		} else if err != nil{
			fmt.Println("Ocorreu um erro", err)
		}
	}
	arquivo.Close()
	return sites
}