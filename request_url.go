package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntroducao()
	fmt.Println("")

	for {
		exibeMenu()
		fmt.Println("")
		comando := leComando()
		fmt.Println("")
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			fmt.Println("")
			imprimeLogs()
		case 3:
			fmt.Println("Obrigado! Até a próxima")
			os.Exit(0) // Saindo sem erro
		default:
			fmt.Println("Comando desconhecido")
			os.Exit(-1) // Saindo com erro
		}
	}
}

func exibeIntroducao() {
	var nome string = "Miguel"
	var versao = 1.2
	idade := 34
	fmt.Println("Olá Sr.", nome, "sua idade é", idade)
	fmt.Printf("Este programa está na versão %.2f\n", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leArquivo()
	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testaSite(site string) {
	resposta, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resposta.StatusCode == 200 {
		fmt.Printf("Site: %s -> Carregado com sucesso\n", site)
		registraLogs(site, true)
	} else {
		fmt.Printf("Site: %s -> Erro no carregamento\n", site)
		registraLogs(site, false)
	}
}

func leArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	defer arquivo.Close()
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}
	return sites
}

func registraLogs(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer arquivo.Close()

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	data := time.Now()
	arquivo.WriteString(fmt.Sprintf("%s - %s - online: %t\n", data.Format(("02/Jan/2006 - 15:04")), site, status))
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}
