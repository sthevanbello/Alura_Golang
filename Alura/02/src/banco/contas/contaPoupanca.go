package contas

import (
	"banco/clientes"
)

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) string {

	if valorSaque <= 0 {
		return "Digite um valor positivo e maior do que zero"
	}
	if valorSaque <= c.saldo {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso!"
	}
	return "saldo insuficiente"
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito <= 0 {
		return "Digite um valor positivo e maior do que zero", c.saldo
	}
	c.saldo += valorDeposito
	return "Depósito efetuado com sucesso", c.saldo
}

func (c *ContaPoupanca) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia > c.saldo || valorTransferencia <= 0 {
		return false
	}
	c.saldo -= valorTransferencia
	contaDestino.Depositar(valorTransferencia)
	return true
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
