package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome string = "Sthevan" // Pode retirar o tipo (string)
	// var idade int = 38          // Pode retirar o tipo (int)
	var versao float32 = 1.1 // Pode retirar o tipo (float), mas nesse caso é bom deixar
	var versaoType = 1.1     // Pode retirar o tipo (float), mas nesse caso é bom deixar

	nomeSemVar := "Nome sem usar o Var" // Pode declarar a variável dessa forma, pois também funciona
	fmt.Println("Olá,", nome)
	// fmt.Println("Idade:", idade)
	fmt.Println("Versão do software:", versao)

	fmt.Println("Tipo da variável versaoType:", reflect.TypeOf(versaoType))

	fmt.Print(nomeSemVar)
}
