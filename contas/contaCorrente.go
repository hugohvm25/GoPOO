// criação de uma estrutura com variáveis que serão usadas com frequencia para não serem repetidas
/*para tornar uma variável publica precisamos tirar ela de qualquer função e colocar a primeira letra maúscula
 */

/*o comando c *ContaCorrente é abreviado para referenciar a estrutura que vai ser utilizada podendo também
 o nome da estrutura, no caso ContaCorrente porém não é a forma comum em Go

 Outra forma de escrita:
 saquePermitido := valorSaque > 0 && valorSaque <= c.Saldo
 if saquePermitido {
	 c.Saldo -= valorSaque

 o && é uma forma de adicionar outra condição para verificação para que seja true ou false e continuar o código
 no caso o valor do saque tem que ser maior que zero ou positivo e ser menor ou igual ao Saldo da conta
*/

// se a condição acima for true, segue, senão pula para o else
// decrementa o valor do saque anterior pelo falor informado ==> c.Saldo = contaHugo.Saldo - valorSaque

/*
 a função deve ter o identificador de qual conta(dado) que deseja buscar. Passa um valor de transferência, e a conta de destino
 pega a referência também pelo identificador.
 faz a verificação para que o valor tranferido seja sempre positivo e menor ou igual ao Saldo da conta obrigatóriamente.
 o valor da transferência é diminuido do Saldo
 a conta de destino recebe o comando de depósito do valor transferido.
*/

package contas

import (
	"example.com/hello/go/src/GoPOO/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular //fazendo link com a struct cliente
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) SacarDinheiro(valorSaque float64) (bool, float64) {
	if valorSaque > 0 && valorSaque <= c.Saldo {
		c.Saldo -= valorSaque
		return true, valorSaque
	} else {
		return false, 0
	}
}

func (c *ContaCorrente) DepositarDinheiro(valorDeposito float64) (bool, float64) {
	if valorDeposito > 0 {
		c.Saldo += valorDeposito
		return true, c.Saldo
	} else {
		return false, c.Saldo
	}
}

func (c *ContaCorrente) Transferencia(valorTransferido float64, contaDestino *ContaCorrente) (bool, float64) {
	if valorTransferido > 0 && valorTransferido <= c.Saldo {
		c.Saldo -= valorTransferido
		statusDeposito, _ := contaDestino.DepositarDinheiro(valorTransferido)
		if statusDeposito {
			return true, valorTransferido
		}
	}
	return false, 0
}
