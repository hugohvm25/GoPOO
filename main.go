/*
PROJETO CURSO ALURA - ORIENTAÇÂO A OBJETOS

Criação de um gerenciador de contas de um banco

## tipo de atribuição de variável simples:

var titular string = "Hugo"
var numeroAgencia int = 555
var numeroConta int = 123456
var saldo float64 = 150.25

ou

## forma reduzida de declaração de variável onde o Go reconhece o tipo sem precisar explicitar

titular := "Hugo"
numeroAgencia := 555
numeroConta := 1234
saldo := 150.25


# Mesmo sem atribuir valor às variáveis, o GO inicializa com as "Zero Initialization"
bool = false
int = 0
float = 0
string = ""
struct = {}


//Criando uma nova conta/variável na estrutura

// sem colocar * o Go não entende de onde puxar a informação; se é da estrutura ou da nova ContaCorrente criada depois
	var contaLaura *ContaCorrente

// criação de um novo bloco de dados para a estrutura
	contaLaura = new(ContaCorrente)

// atribuição dos dados exigidos pela estrutura para a Nova ContaCorrente
	contaLaura.titular = "Laura"
	contaLaura.numeroAgencia = 222

// irá gerar no console a informação SEM rotulação porém com os dados informados
	fmt.Println(contaLaura)
	// console // &{Laura 222 0 0} => onde o & é um endereço

// irá gerar no console a informação COM rotulação (nome atribuido) com os dados informados
	fmt.Println(*contaLaura)
	// console // {Laura 222 0 0}


# metodo de comparação e verificação se os 2 blocos de dados (conteúdo) são iguais ou diferentes

	contaHugo := ContaCorrente{titular: "Hugo", numeroAgencia: 555, numeroConta: 1234, saldo: 150.25}
	contaHugo2 := ContaCorrente{titular: "Hugo", numeroAgencia: 555, numeroConta: 1234, saldo: 150.25}

	fmt.Println(contaHugo == contaHugo2)
	// console // true (retorna um boolenado informando que é verdandeiro apesar de serem contas diferentes pois possuem os mesmos dados)



	contaHugo := contaCorrente{titular: "Hugo", numeroAgencia: 555, numeroConta: 1234, saldo: 150.25}
	fmt.PrintLn(contaHugo.saldo)
	// console // 150.25

	valorSaque = 50
	contaHugo.saldo = contaHugo.saldo - valorSaque
	fmt.Println(contaHugo.saldo)
	// console // 100.25

*/

package main

import "fmt"

//criação de uma estrutura com variáveis que serão usadas com frequencia para não serem repetidas
type contaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	contaHugo := contaCorrente{titular: "Hugo", numeroAgencia: 555, numeroConta: 1234, saldo: 150.10} // forma extensa de passar os dados para a estrutura
	//contaIsabela := ContaCorrente{"Isabela", 500, 4321, 100.00}                                       // forma curta de passar dados para a estrutura

	fmt.Println(contaHugo)
	//fmt.Println(contaIsabela)

	fmt.Println(contaHugo.saldo)
	fmt.Println(contaHugo.sacarDinheiro(50))
	fmt.Println(contaHugo.saldo)

	fmt.Println(contaHugo.depositarDinheiro(200))
	fmt.Println(contaHugo.saldo)
}

func (c *contaCorrente) sacarDinheiro(valorSaque float64) string {
	// o comando c *contaCorrente é abreviado para referenciar a estrutura que vai ser utilizada podendo também
	// o nome da estrutura, no caso contaCorrente porém não é a forma comum em Go
	saquePermitido := valorSaque > 0 && valorSaque <= c.saldo
	// o && é uma forma de adicionar outra condição para verificação para que seja true ou false e continuar o código
	// no caso o valor do saque tem que ser maior que zero ou positivo e ser menor ou igual ao saldo da conta

	if saquePermitido { // se a condição acima for true, segue, senão pula para o else
		c.saldo -= valorSaque // decrementa o valor do saque anterior pelo falor informado ==> c.saldo = contaHugo.saldo - valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *contaCorrente) depositarDinheiro(valorDeposito float64) string {

	depositoPermitido := valorDeposito > 0
	if depositoPermitido {
		c.saldo += valorDeposito
		return "Depósito realizado com sucesso"
	} else {
		return "Depósito negado"
	}

}
