package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

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

	fmt.Println()
	fmt.Println("====== Monitor de sites ======")
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
		iniciarMonitoramento()
	case 2:
		fmt.Println("Opção Selecionada: Exibir Log de Monitoramento")
	default:
		exibirOpcaoInvalida()
	}
}

func sairDoPrograma() {
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
	fmt.Println("Opção Selecionada: Iniciar Monitoramento")
	site := "https://www.alura.com.br"
	siteBinNotFound := "https://httpbin.org/status/404"
	siteBinOk := "https://httpbin.org/status/200"
	// response, erro := http.Get(site) // Para ignorar um dos objetos retornados, podemos utilizar o operador null _
	response, _ := http.Get(site)
	responseNotFound, _ := http.Get(siteBinNotFound)
	responseOk, _ := http.Get(siteBinOk)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(response.Header)
	fmt.Println(response.ContentLength)
	fmt.Println(response.Request.Method)
	fmt.Println(response.Request.RequestURI)
	fmt.Println("==============================")
	// fmt.Println(erro)

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, " está funcionando corretamente")
	} else if response.StatusCode == 404 {
		fmt.Println("Rota não encontrada:", site, " - Status: ", response.StatusCode)
	}

	if responseNotFound.StatusCode == 404 {
		fmt.Println("Response 404")
	}

	if responseOk.StatusCode == 200 {
		fmt.Println("Response OK 200")
	} else {
		fmt.Println(responseOk.Body)
	}
}
