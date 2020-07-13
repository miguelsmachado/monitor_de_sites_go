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

	for { 
		exibeMenu()

		comando := leComando()
		fmt.Println("")

		switch comando {
		case 1:
			iniciaMonitoramento()
		case 2:
			exibeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}

	}

}

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

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")
	fmt.Println("")

	sites := []string{"https://www.alura.com.br", "https://www.caelum.com.br"}
	sites = append(sites, "https://random-status-code.herokuapp.com")

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Site nr", i+1, ":", site)
			testanodSite(site)
		}

		monitoramentoComArquivo()

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func testanodSite(site string) {
	resp, _ := http.Get(site) 

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