package main

import "fmt"

func main() {

	var titular string = "Sthevan" // titular := "Sthevan"
	var numeroAgencia int = 589
	var numeroConta int = 456789
	var saldo float64 = 125.50

	contaCorrente := ContaCorrente{
		titular:       titular,
		numeroAgencia: numeroAgencia,
		numeroConta:   numeroConta,
		saldo:         saldo}

	var contaNova *ContaCorrente
	contaNova = new(ContaCorrente)

	contaNova.titular = "Homer"
	contaNova.numeroAgencia = 123
	contaNova.numeroConta = 456123
	contaNova.saldo = 850.55

	fmt.Println(contaCorrente)
	fmt.Println(*contaNova)

}

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}
