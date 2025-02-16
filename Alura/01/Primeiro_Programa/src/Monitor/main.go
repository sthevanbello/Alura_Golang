package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
)

const timeLoops = 5
const timeDelay = 5

func main() {
	clearScreen()
	nome := "Sthevan"
	versao := 1.1
	exibirIntroducao(nome, versao)
	for {
		exibirMenu()
		opcao := lerOpcaoMenu()
		exibirOpcoes(opcao)
	}
}

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func exibirIntroducao(nome string, versao float64) {
	tituloColorido := fmt.Sprintf(color.RedString("====== Monitor de sites ======"))
	fmt.Println(tituloColorido)
	fmt.Println()
	fmt.Println("Bem vindo,", nome)
	fmt.Println("Versão do monitor:", versao)
	fmt.Println("==============================")
	fmt.Println()
}

func exibirMenu() {
	fmt.Println("----- Menu Inicial -----")
	fmt.Println()
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Log de Monitoramento")
	fmt.Println("0 - Sair do Monitoramento")
	fmt.Println("==============================")
}

func lerOpcaoMenu() int {
	var opcaoLida int
	fmt.Print("Escolha uma opção do menu: ")
	fmt.Scanf("%d\n", &opcaoLida)
	// fmt.Scan(&opcao) // Pode ser utilizada a function fmt.Scan ao invés de fmt.Scanf para não precisar colocar o tipo de entrada %d para int, por exemplo.
	return opcaoLida
}

func exibirOpcoes(opcao int) {

	// if opcao == 0 {
	// 	fmt.Println("Opção selecionada: Sair")
	// } else if opcao == 1 {
	// 	fmt.Println("Opção Selecionada: Iniciar Monitoramento")
	// } else if opcao == 2 {
	// 	fmt.Println("Opção Selecionada: Exibir Log de Monitoramento")
	// } else {
	// 	fmt.Println("Opção Inválida")
	// }

	switch opcao {
	case 0:
		sairDoPrograma()
	case 1:
		// iniciarMonitoramento()
		// iniciarMonitoramentoSlice()
		monitorarSitesFile()
	case 2:
		fmt.Println("Opção Selecionada: Exibir Log de Monitoramento")
	default:
		exibirOpcaoInvalida()
	}
}

func sairDoPrograma() {
	clearScreen()
	fmt.Println()
	fmt.Println("==============================")
	fmt.Println("Opção selecionada: Sair")
	fmt.Println("==============================")
	fmt.Println("Saindo...")
	os.Exit(0)
}

func exibirOpcaoInvalida() {
	fmt.Println()
	fmt.Println("==============================")
	fmt.Println("Opção Inválida")
	fmt.Println("==============================")
	os.Exit(-1)
}

func iniciarMonitoramento() {
	fmt.Println("Opção Selecionada: Iniciar Monitoramento utilizando Array")
	var sites [4]string
	sites[0] = "https://www.alura.com.br"
	sites[1] = "https://httpbin.org/status/404"
	sites[2] = "https://httpbin.org/status/200"

	monitorarSitesArray(sites)
}

func iniciarMonitoramentoSlice() {
	fmt.Println("Opção Selecionada: Iniciar Monitoramento utilizando Slice")
	sites := []string{
		"https://www.alura.com.br",
		"https://httpbin.org/status/404",
		"https://httpbin.org/status/200",
	}
	// fmt.Println(sites)
	sites = append(sites, "https://www.meutimao.com.br")
	// fmt.Println(sites)
	monitorarSitesSlice(sites)
}

func monitorarSitesArray(sites [4]string) {
	fmt.Println()
	for i := 0; i < len(sites); i++ {
		response, _ := http.Get(sites[i])

		if response.StatusCode == 200 {
			fmt.Println("Site:", sites[i], " está funcionando corretamente")
		} else if response.StatusCode == 404 {
			fmt.Println("Rota não encontrada:", sites[i], " - Status: ", response.StatusCode)
		}
	}
	fmt.Println()
	fmt.Println("==============================")
}

func monitorarSitesSlice(sites []string) {
	fmt.Println()
	fmt.Println(sites)
	fmt.Println(len(sites))
	// for i := 0; i < len(sites); i++ {
	// 	response, _ := http.Get(sites[i])

	// 	if response.StatusCode == 200 {
	// 		fmt.Println("Site:", sites[i], " está funcionando corretamente")
	// 	} else if response.StatusCode == 404 {
	// 		fmt.Println("Rota não encontrada:", sites[i], " - Status: ", response.StatusCode)
	// 	} else {
	// 		fmt.Println("O site:", sites[i], "Não está respondendo. Status Code:", response.StatusCode)
	// 	}
	// }
	for i := 0; i < timeLoops; i++ {
		for i, site := range sites {
			fmt.Println("Site: ", i+1)
			testarSite(site)
		}
		time.Sleep(timeDelay * time.Second)
		fmt.Println("---------------------------------------------------------------------------------------------")
	}
	fmt.Println()
}

func monitorarSitesFile() {
	sites := lerArquivoSites()
	fmt.Println()
	for i := 0; i < timeLoops; i++ {
		tituloColorido := fmt.Sprintf(color.CyanString("--------------------------- Monitorando Sites do Arquivo -----------------------------------"))
		fmt.Println(tituloColorido)

		colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, "Hello")
		fmt.Println(colored)

		for i, site := range sites {
			fmt.Println("Site: ", i+1)
			testarSite(site)
		}
		time.Sleep(timeDelay * time.Second)
	}
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Println()
}

func testarSite(site string) {
	response, err := http.Get(site)
	if err != nil {
		fmt.Println(err)
	} else if response.StatusCode == 200 {
		fmt.Println("Site:", site, " está funcionando corretamente - Status Code:", response.StatusCode)
	} else if response.StatusCode == 404 {
		fmt.Println("Rota não encontrada:", site, " - Status Code: ", response.StatusCode)
	} else {
		fmt.Println("O site:", site, "Não está respondendo. Status Code:", response.StatusCode)
	}
	registrarLog(site, response.StatusCode)
}

func lerArquivoSites() []string {
	clearScreen()
	var retorno []string
	// arquivo, err := os.Open("sites.txt")
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		// fmt.Println(string(arquivo))
		leitor := bufio.NewReader(arquivo)
		for {
			linha, err := leitor.ReadString('\n')
			fmt.Println("--------------------------- Monitorando Sites do Arquivo -----------------------------------")
			fmt.Println(linha)
			linha = strings.TrimSpace(linha)
			retorno = append(retorno, linha)
			if err == io.EOF {
				// fmt.Println(err)
				break
			}

		}
	}
	arquivo.Close()
	return retorno
}

func registrarLog(site string, status int) {
	var tituloColorido string
	var mensagem string
	if status != 200 {
		tituloColorido = fmt.Sprintf(color.RedString("==================================== Registro de Log - Error ===================================="))
		mensagem = "Erro"
	} else {
		tituloColorido = fmt.Sprintf(color.GreenString("==================================== Registro de Log - Ok ===================================="))
		mensagem = "Sucesso"
	}

	linha := fmt.Sprintf("Site: %-50sStatus: %-10d Retorno: %s\n", site, status, mensagem)
	fmt.Println(tituloColorido)

	fmt.Println(site, status)

	fimColorido := fmt.Sprintf(color.HiBlueString("==================================== Registro de Log - Fim ===================================="))
	fmt.Println(fimColorido)

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(linha)
	fmt.Println(linha)
}
