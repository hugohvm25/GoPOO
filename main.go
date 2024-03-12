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

	contaHugo := contaCorrente{titular: "Hugo", numeroAgencia: 555, numeroConta: 1234, saldo: 500} // forma extensa de passar os dados para a estrutura
	contaIsabela := contaCorrente{"Isabela", 500, 4321, 100.00}                                    // forma curta de passar dados para a estrutura

	fmt.Println("Dados das Contas:")
	fmt.Println(contaHugo)
	fmt.Println(contaIsabela)
	fmt.Println("")

	fmt.Println("Conta do:", contaHugo.titular, "-", "Saldo atualizado:", contaHugo.saldo)
	fmt.Println("Conta da:", contaIsabela.titular, "-", "Saldo atualizado:", contaIsabela.saldo)
	fmt.Println("")

	statusSaque, valor2 := contaHugo.sacarDinheiro(100)
	fmt.Println(contaHugo.titular, "sacou o valor:", valor2, "-", statusSaque)
	fmt.Println("Conta do:", contaHugo.titular, "-", "Saldo atualizado:", contaHugo.saldo)
	fmt.Println("")

	statusDeposito, valor := contaHugo.depositarDinheiro(300)
	fmt.Println(contaHugo.titular, "depositou o valor:", valor, "-", statusDeposito)
	fmt.Println("Conta do:", contaHugo.titular, "-", "Saldo atualizado:", contaHugo.saldo)
	fmt.Println("")

	statusTransf, valorTransf := contaHugo.transferencia(100, &contaIsabela) // conta de origem da transação, comando(valor da transf, conta destino com identificador &)
	// para identificar para qual conta vai ser transferido o valor, é preciso usar o & que vai direcionar ao endereço daquela conta em específico.
	fmt.Println("A transferência ocorreu?", statusTransf)
	fmt.Println("O valor tranferido de", contaHugo.titular, "para", contaIsabela.titular, "foi:", valorTransf)
	fmt.Println("Conta do:", contaHugo.titular, "-", "Saldo atualizado:", contaHugo.saldo)
	fmt.Println("Conta da:", contaIsabela.titular, "-", "Saldo atualizado:", contaIsabela.saldo)
	fmt.Println("")

}

func (c *contaCorrente) sacarDinheiro(valorSaque float64) (string, float64) {
	/*o comando c *contaCorrente é abreviado para referenciar a estrutura que vai ser utilizada podendo também
	o nome da estrutura, no caso contaCorrente porém não é a forma comum em Go

	Outra forma de escrita:
	saquePermitido := valorSaque > 0 && valorSaque <= c.saldo
	if saquePermitido {
		c.saldo -= valorSaque

	o && é uma forma de adicionar outra condição para verificação para que seja true ou false e continuar o código
	no caso o valor do saque tem que ser maior que zero ou positivo e ser menor ou igual ao saldo da conta
	*/

	if valorSaque > 0 && valorSaque <= c.saldo { // se a condição acima for true, segue, senão pula para o else
		c.saldo -= valorSaque // decrementa o valor do saque anterior pelo falor informado ==> c.saldo = contaHugo.saldo - valorSaque
		return "Saque realizado com sucesso", valorSaque
	} else {
		return "Saldo insuficiente", valorSaque
	}
}

func (c *contaCorrente) depositarDinheiro(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Depósito realizado com sucesso", valorDeposito
	} else {
		return "Depósito negado", valorDeposito
	}
}

func (c *contaCorrente) transferencia(valorTransferido float64, contaDestino *contaCorrente) (bool, float64) {
	if valorTransferido > 0 && valorTransferido <= c.saldo {
		c.saldo -= valorTransferido
		contaDestino.depositarDinheiro(valorTransferido)
		return true, valorTransferido
	} else {
		return false, valorTransferido
	}
	/*
		a função deve ter o identificador de qual conta(dado) que deseja buscar. Passa um valor de transferência, e a conta de destino
		pega a referência também pelo identificador.
		faz a verificação para que o valor tranferido seja sempre positivo e menor ou igual ao saldo da conta obrigatóriamente.
		o valor da transferência é diminuido do saldo
		a conta de destino recebe o comando de depósito do valor transferido.
	*/

}
