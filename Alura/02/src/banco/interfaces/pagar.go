package interfaces

type tipoConta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta tipoConta, valor float64) {
	conta.Sacar(valor)
}
