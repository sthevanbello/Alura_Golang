package contas

import (
	"banco/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {

	if valorSaque <= 0 {
		return "Digite um valor positivo e maior do que zero"
	}
	if valorSaque <= c.Saldo {
		c.Saldo -= valorSaque
		return "Saque realizado com sucesso!"
	}
	return "Saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito <= 0 {
		return "Digite um valor positivo e maior do que zero", c.Saldo
	}
	c.Saldo += valorDeposito
	return "Depósito efetuado com sucesso", c.Saldo
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia > c.Saldo || valorTransferencia <= 0 {
		return false
	}
	c.Saldo -= valorTransferencia
	contaDestino.Depositar(valorTransferencia)
	return true
}
