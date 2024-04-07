package contas

import "example.com/hello/go/src/GoPOO/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int // forma linear de declarar as variáveis do mesmo tipo de dado
	saldo                                float64
}

func (c *ContaPoupanca) SacarDinheiro(valorSaque float64) string {
	if valorSaque > 0 && valorSaque <= c.saldo {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaPoupanca) DepositarDinheiro(valorDeposito float64) (bool, string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return true, "Deposito realizado com sucesso", c.saldo
	} else {
		return false, "Não foi possivel realizar o deposito", c.saldo
	}
}

func (c *ContaPoupanca) ConsultaSaldo() float64 {
	return c.saldo
}
