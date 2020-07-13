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
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		lista = append(lista, linha)

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Ocorreu um erro:", err)
		}
	}
	arquivo.Close()

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
		arquivo.WriteString(time.Now().Format("02/01/2006 - 15:04:05 -> ") + site + " esta online\n")
	} else if status == false{
		arquivo.WriteString(time.Now().Format("02/01/2006 - 15:04:05 -> ") + site + " esta offline\n")
	}

	arquivo.Close()
}

func exibeLogs(){
	fmt.Println("Exibindo logs...")
	fmt.Println("")

	arquivo, _  := ioutil.ReadFile("log.txt")

	fmt.Println(string(arquivo))
}