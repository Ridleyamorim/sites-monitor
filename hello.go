package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 30

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs ...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao()  {
	name := "Ridley"
	version := 1.1

	fmt.Println("Olá, Sr.", name)
	fmt.Println("A sua versão é a", version)
}

func exibeMenu()  {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir os logs")
	fmt.Println("0 - Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento()  {
	fmt.Println("Monitorando...")

	sites := leSitesdoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for  _, site := range sites{
			fmt.Println("Site:", site)
			testaSite(site)
			fmt.Println("")
		}
		time.Sleep(delay * time.Minute)
		fmt.Println("")
	}
	
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	switch resp.StatusCode {
		case 200:
			fmt.Println("Foi carregado com sucesso, Status Code:", resp.StatusCode)
		case 401:
			fmt.Println("A solicitação não foi aplicada porque não possui credenciais de autenticação válidas para o recurso de destino, Status Code:", resp.StatusCode)
		case 403:
			fmt.Println("O servidor não autorizou a emissão de um resposta, Status Code:", resp.StatusCode)
		case 404:
			fmt.Println("Não foi encontrado, Status Code:", resp.StatusCode)
		default:
			fmt.Println("Houve problemas, Status Code:", resp.StatusCode)
	}
}

func leSitesdoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF{
			break
		}
	}

	arquivo.Close()

	return sites
}