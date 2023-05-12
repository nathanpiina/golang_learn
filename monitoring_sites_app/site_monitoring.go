package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 1
const delay = 5

func main(){

	exibeIntroducao()
	leSitesDoArquivo()
	exibeMenu()
	comando := lerComando()

	switch comando {
		case 1:
			iniciarMonitoramento()

		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()

		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)

		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
	}

}

func exibeIntroducao(){
	var nome string = "Nathan"
	var versao float32 = 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Esse programa está na versão", versao)
}

func lerComando() int{
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func exibeMenu(){
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}

func iniciarMonitoramento(){

	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()

	for i := 0; i <= monitoramentos; i++ {

		for i, site := range sites{
			fmt.Println("Estou passando na posição", i, "do meu slice e a posicao tem o site:", site)
			testaSite(site)
		}

		time.Sleep((5 * time.Second))
		fmt.Println("")

	}

}

func testaSite(site string) {

	resp, err := http.Get(site)

		if err != nil{
			fmt.Println("Ocorreu um erro:", err)
		}

		if resp.StatusCode == 200{
			fmt.Println("Site:", site, "foi carregado com sucesso!")
			registraLogs(site, true)

		}else{
			fmt.Println("Site:", site, "esta com problemas, Status code:", resp.StatusCode )
			registraLogs(site, false)
		}
}

func leSitesDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil{
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

	arquivo.Close()

	return sites
}

func registraLogs(site string, status bool){

	arquivo, err := os.OpenFile("logS.txt", os.O_CREATE|os.O_RDWR | os.O_APPEND, 0666)

	if err != nil{
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- Online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()

}

func imprimeLogs(){

	arquivo, err := ioutil.ReadFile("logs.txt")

	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}