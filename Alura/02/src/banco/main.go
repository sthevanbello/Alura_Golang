package main

import (
	"banco/contas"
	"fmt"
)

func main() {

	var titular string = "Homer" // Titular := "Sthevan"
	var numeroAgencia int = 589
	var numeroConta int = 456789
	var saldo float64 = 825.50

	contaCorrentePrimaria := contas.ContaCorrente{
		Titular:       titular,
		NumeroAgencia: numeroAgencia,
		NumeroConta:   numeroConta,
		Saldo:         saldo}

	contaCorrenteSecundaria := contas.ContaCorrente{Titular: "Marge", NumeroAgencia: 589, NumeroConta: 159753, Saldo: 1500}

	fmt.Println(contaCorrentePrimaria)
	fmt.Println(contaCorrenteSecundaria)

	status := contaCorrentePrimaria.Transferir(-100, &contaCorrenteSecundaria)
	fmt.Println(status)

	fmt.Println(contaCorrentePrimaria)
	fmt.Println(contaCorrenteSecundaria)
}
