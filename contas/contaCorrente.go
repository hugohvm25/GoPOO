// criação de uma estrutura com variáveis que serão usadas com frequencia para não serem repetidas
/*para tornar uma variável publica precisamos tirar ela de qualquer função e colocar a primeira letra maúscula
 */

/*o comando c *ContaCorrente é abreviado para referenciar a estrutura que vai ser utilizada podendo também
 o nome da estrutura, no caso ContaCorrente porém não é a forma comum em Go

 Outra forma de escrita:
 saquePermitido := valorSaque > 0 && valorSaque <= c.saldo
 if saquePermitido {
	 c.saldo -= valorSaque

 o && é uma forma de adicionar outra condição para verificação para que seja true ou false e continuar o código
 no caso o valor do saque tem que ser maior que zero ou positivo e ser menor ou igual ao saldo da conta
*/

// se a condição acima for true, segue, senão pula para o else
// decrementa o valor do saque anterior pelo falor informado ==> c.saldo = contaHugo.saldo - valorSaque

/*
 a função deve ter o identificador de qual conta(dado) que deseja buscar. Passa um valor de transferência, e a conta de destino
 pega a referência também pelo identificador.
 faz a verificação para que o valor tranferido seja sempre positivo e menor ou igual ao saldo da conta obrigatóriamente.
 o valor da transferência é diminuido do saldo
 a conta de destino recebe o comando de depósito do valor transferido.


 O saldo está publico e visivel para todos podendo ser alterado, deve ser feita a troca para que seja privado mas será feita uma função para consutlar o saldo.
*/

package contas

import (
	"example.com/hello/go/src/GoPOO/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular //fazendo link com a struct cliente
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func NovaContaCorrente(titular clientes.Titular, numeroAgencia int, numeroConta int, saldo float64) (c *ContaCorrente) {
	return &ContaCorrente{
		Titular:       titular,
		NumeroAgencia: numeroAgencia,
		NumeroConta:   numeroConta,
		saldo:         saldo,
	}
}

func (c *ContaCorrente) SacarDinheiro(valorSaque float64) (bool, float64) {
	if valorSaque > 0 && valorSaque <= c.saldo {
		c.saldo -= valorSaque
		return true, valorSaque
	} else {
		return false, 0
	}
}

func (c *ContaCorrente) DepositarDinheiro(valorDeposito float64) (bool, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return true, c.saldo
	} else {
		return false, c.saldo
	}
}

func (c *ContaCorrente) Transferencia(valorTransferido float64, contaDestino *ContaCorrente) (bool, float64) {
	if valorTransferido > 0 && valorTransferido <= c.saldo {
		c.saldo -= valorTransferido
		statusDeposito, _ := contaDestino.DepositarDinheiro(valorTransferido)
		if statusDeposito {
			return true, valorTransferido
		}
	}
	return false, 0
}

func (c *ContaCorrente) ConsultaSaldo() float64 {
	return c.saldo
}
