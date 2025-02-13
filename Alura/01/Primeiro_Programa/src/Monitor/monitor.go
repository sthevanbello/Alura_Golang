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
	fmt.Println("2 - Exibir Log de Monitoramento")
	fmt.Println("0 - Sair do Monitoramento")
	fmt.Println("==============================")

	fmt.Print("Escolha uma opção do menu: ")
	fmt.Scanf("%d", &opcao)
	// fmt.Scan(&opcao) // Pode ser utilizada a function fmt.Scan ao invés de fmt.Scanf para não precisar colocar o tipo de entrada %d para int, por exemplo.

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
		fmt.Println("Opção selecionada: Sair")
	case 1:
		fmt.Println("Opção Selecionada: Iniciar Monitoramento")
	case 2:
		fmt.Println("Opção Selecionada: Exibir Log de Monitoramento")
	default:
		fmt.Println("Opção Inválida")
	}
}
