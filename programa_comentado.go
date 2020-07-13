package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	"io/ioutil"
)

const monitoramento = 3
const delay = 5

func main() {

	exibeIntroducao()

	for { //criar um for sem condição nenhuma é igual a um while true.
		exibeMenu()

		comando := leComando()
		fmt.Println("")

		// Usando switch (in case of do pascal)
		switch comando {
		case 1:
			iniciaMonitoramento()
		case 2:
			exibeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0) // Saindo do programa sem erros
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1) // Saindo do programa com erro
		}

	}

}

// FUNÇÕES EM GO

//Função sem entrada e sem retorno
func exibeIntroducao() {
	var nome string = "Miguel"
	versao := 1.1
	fmt.Println("Olá sr.", nome)
	fmt.Println("Esse programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("\n1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir os logs")
	fmt.Println("0 - Sair do programa")
}

//Função sem entrada e com retorno
//Devemos declarar ao lado do nome da função qual o tipo do retorno.
func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")

	// Declarando um array:
	// var <nome> [tamanho] <tipo>
	// var sites [4]string
	// incluindo dados em um array
	//sites[0] = "https://www.alura.com.br"

	// Utilizando slices
	sites := []string{"https://www.alura.com.br", "https://www.caelum.com.br"}
	// Realizando um append
	sites = append(sites, "https://random-status-code.herokuapp.com")

	// utilizando o for para percorrer o slice:
	// for normal:
	// for i:=0; i < len(sites); i++ {

	// for otimizado para GO, usando o range
	// for <indice>, <conteudo> := range <array>
	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Site nr", i+1, ":", site)
			testanodSite(site)
		}

		// Para fins didaticos, fiz esse outro monitoramento utilizando a abertura de um arquivo
		monitoramentoComArquivo()

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func testanodSite(site string) {
	resp, _ := http.Get(site) // requisição get em http
	// fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("O site", site, "está fora do ar!")
		registraLog(site, false)
	}
}

func monitoramentoComArquivo() {
	var lista []string

	// Abrimos o arquivo com o os
	arquivo, err := os.Open("sites.txt")
	// tratando possíveis erros
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

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
		lista = append(lista, linha)

		// vamos usar o erro EOF (End of File) para encerrar o loop quando não houver mais nada para ser lido
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Ocorreu um erro:", err)
		}
	}

	arquivo.Close()

	// Realizando o monitoramento
	for i := 0; i < monitoramento; i++ {
		for i, site := range lista {
			fmt.Println("Site nr", i+1, ":", site)
			testanodSite(site)
		}
	}
}

func registraLog(site string, status bool) {

	// O método OpenFile, permite que um arquivo seja aberto de uma forma mais detalhada.
	// Além do nome do arquivo, passamos algumas flags para dizer como queremos trabalhar
	// Nesse caso, usamos os.O_RDWR (abre no modo leitura e escrita) e os.O_CREATE(se o arquivo não
	// existir, ele o cria) e uma permissão (usamos 0666)
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil{
		fmt.Println("Ocorreu um erro", err)
	}
	if status == true {
		// O go utiliza um padrao de escrita para data/hora, consultar documentação antes de preencher
		arquivo.WriteString(time.Now().Format("02/01/2006 - 15:04:05 -> ") + site + " esta online\n")
	} else if status == false{
		arquivo.WriteString(time.Now().Format("02/01/2006 - 15:04:05 -> ") + site + " esta offline\n")
	}

	arquivo.Close()
}

func exibeLogs(){
	fmt.Println("Exibindo logs...")
	fmt.Println("")

	//A biblioteca io/ioutil permite que se leia o arquivo todo de uma só vez
	// Não é necessário fechar.
	arquivo, _  := ioutil.ReadFile("log.txt")

	fmt.Println(string(arquivo))
}