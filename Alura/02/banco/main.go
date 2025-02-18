package main

import "fmt"

func main() {

	var titular string = "Homer" // titular := "Sthevan"
	var numeroAgencia int = 589
	var numeroConta int = 456789
	var saldo float64 = 825.50

	contaCorrentePrimaria := ContaCorrente{
		titular:       titular,
		numeroAgencia: numeroAgencia,
		numeroConta:   numeroConta,
		saldo:         saldo}

	contaCorrenteSecundaria := ContaCorrente{titular: "Marge", numeroAgencia: 589, numeroConta: 159753, saldo: 1500}

	fmt.Println(contaCorrentePrimaria)
	fmt.Println(contaCorrenteSecundaria)

	status := contaCorrentePrimaria.Transferir(-100, &contaCorrenteSecundaria)
	fmt.Println(status)

	fmt.Println(contaCorrentePrimaria)
	fmt.Println(contaCorrenteSecundaria)
}

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {

	if valorSaque <= 0 {
		return "Digite um valor positivo e maior do que zero"
	}
	if valorSaque <= c.saldo {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso!"
	}
	return "Saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito <= 0 {
		return "Digite um valor positivo e maior do que zero", c.saldo
	}
	c.saldo += valorDeposito
	return "Depósito efetuado com sucesso", c.saldo
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia > c.saldo || valorTransferencia <= 0 {
		return false
	}
	c.saldo -= valorTransferencia
	contaDestino.Depositar(valorTransferencia)
	return true
}
