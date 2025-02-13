package main

import "fmt"

func main() {
	nome := "Sthevan"
	versao := 1.1
	var opcao int
	fmt.Println()
	fmt.Println("====== Monitor de sites ======")
	fmt.Println("Bem vindo,", nome)
	fmt.Println("Versão do monitor:", versao)
	fmt.Println("==============================")
	fmt.Println()
	fmt.Println("----- Menu Inicial -----")
	fmt.Println()
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Log Monitoramento")
	fmt.Println("0 - Sair do Monitoramento")
	fmt.Println("==============================")

	fmt.Print("Escolha uma opção do menu: ")
	fmt.Scanf("%d", &opcao)
	// fmt.Scan(&opcao) // Pode ser utilizada a function fmt.Scan ao invés de fmt.Scanf para não precisar colocar o tipo de entrada %d para int, por exemplo.

	fmt.Println("Opção selecionada:", opcao)
}
