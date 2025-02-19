package main

import (
	"banco/clientes"
	"banco/contas"
	"banco/interfaces"
	"fmt"
)

func main() {

	titular := clientes.Titular{Nome: "Homer", CPF: "97201075888", Profissao: "Operador de usina nuclear"}
	var numeroAgencia int = 589
	var numeroConta int = 456789

	contaCorrentePrimaria := contas.ContaCorrente{
		Titular:       titular,
		NumeroAgencia: numeroAgencia,
		NumeroConta:   numeroConta}

	contaCorrenteSecundaria := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Marge",
			CPF:       "32318907803",
			Profissao: "Dona de casa"},
		NumeroAgencia: 589,
		NumeroConta:   159753}

	contaPoupancaPrimaria := contas.ContaPoupanca{
		Titular: clientes.Titular{
			"Homer",
			"97201075888",
			"Operador de usina nuclear",
		},
		NumeroAgencia: 589,
		NumeroConta:   312456,
		Operacao:      1}

	contaCorrentePrimaria.Depositar(500)
	contaPoupancaPrimaria.Depositar(5000)
	contaCorrenteSecundaria.Depositar(800)
	fmt.Println("Corrente Primária:", contaCorrentePrimaria)
	fmt.Println("Poupança Primária", contaPoupancaPrimaria)
	fmt.Println(contaCorrenteSecundaria)

	status := contaCorrentePrimaria.Transferir(150, &contaCorrenteSecundaria)
	fmt.Println(status)

	fmt.Println(contaCorrentePrimaria)
	fmt.Println(contaCorrenteSecundaria)
	fmt.Println(contaCorrentePrimaria.Obtersaldo())

	interfaces.PagarBoleto(&contaPoupancaPrimaria, 4829)
	fmt.Println("Poupança Primária", contaPoupancaPrimaria.ObterSaldo())
}
