package main

import "fmt"

func main() {

	var titular string = "Homer" // titular := "Sthevan"
	var numeroAgencia int = 589
	var numeroConta int = 456789
	var saldo float64 = 825.50

	contaCorrente := ContaCorrente{
		titular:       titular,
		numeroAgencia: numeroAgencia,
		numeroConta:   numeroConta,
		saldo:         saldo}

	fmt.Println(contaCorrente)
	fmt.Println(contaCorrente.Sacar(-10))
	fmt.Println(contaCorrente)
	fmt.Println(contaCorrente.Depositar(400))
	status, saldo := contaCorrente.Depositar(150)
	fmt.Println(status, saldo)
	fmt.Println(contaCorrente)
	fmt.Println(contaCorrente.Depositar(0))
	fmt.Println(contaCorrente)

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
