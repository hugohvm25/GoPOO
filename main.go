/*
PROJETO CURSO ALURA - ORIENTAÇÂO A OBJETOS

Criação de um gerenciador de contas de um banco

## tipo de atribuição de variável simples:

var Titular string = "Hugo"
var numeroAgencia int = 555
var numeroConta int = 123456
var ConsultaSaldo float64 = 150.25

ou

## forma reduzida de declaração de variável onde o Go reconhece o tipo sem precisar explicitar

Titular := "Hugo"
numeroAgencia := 555
numeroConta := 1234
ConsultaSaldo := 150.25


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
	contaLaura.Titular = "Laura"
	contaLaura.numeroAgencia = 222

// irá gerar no console a informação SEM rotulação porém com os dados informados
	fmt.Println(contaLaura)
	// console // &{Laura 222 0 0} => onde o & é um endereço

// irá gerar no console a informação COM rotulação (nome atribuido) com os dados informados
	fmt.Println(*contaLaura)
	// console // {Laura 222 0 0}


# metodo de comparação e verificação se os 2 blocos de dados (conteúdo) são iguais ou diferentes

	contaHugo := ContaCorrente{Titular: "Hugo", numeroAgencia: 555, numeroConta: 1234, ConsultaSaldo: 150.25}
	contaHugo2 := ContaCorrente{Titular: "Hugo", numeroAgencia: 555, numeroConta: 1234, ConsultaSaldo: 150.25}

	fmt.Println(contaHugo == contaHugo2)
	// console // true (retorna um boolenado informando que é verdandeiro apesar de serem contas diferentes pois possuem os mesmos dados)



	contaHugo := ContaCorrente{Titular: "Hugo", numeroAgencia: 555, numeroConta: 1234, ConsultaSaldo: 150.25}
	fmt.PrintLn(contaHugo.ConsultaSaldo)
	// console // 150.25

	valorSaque = 50
	contaHugo.ConsultaSaldo = contaHugo.ConsultaSaldo - valorSaque
	fmt.Println(contaHugo.ConsultaSaldo)
	// console // 100.25


	[Info  - 12:28:25] 2024/03/24 12:28:25 go/packages.Load #188
	snapshot=1487
	directory=file:///C:/Users/Hugo/go/src/GoPOO

	package="example.com/hello/go/src/GoPOO"



	files=[C:\Users\Hugo\go\src\GoPOO\main.go]





	clienteIsabela := clientes.Titular{"Isabela", "111.222.333.44"}
	contaIsabela := contas.NovaContaCorrente(clienteIsabela, 500, 4321, 300)

	//contaHugo := contas.ContaCorrente{Titular: "Hugo", NumeroAgencia: 555, NumeroConta: 1234, ConsultaSaldo: 500}
	// contaIsabela := contas.ContaCorrente{Titular: "Isabela", NumeroAgencia: 500, NumeroConta: 4321, ConsultaSaldo: 100}

	fmt.Println("Dados das Contas:")
	fmt.Println(contaHugo)
	fmt.Println(contaIsabela)
	fmt.Println("")

	fmt.Println("Conta do:", contaHugo.Titular, "-", "ConsultaSaldo atualizado:", contaHugo.ConsultaSaldo())
	fmt.Println("Conta da:", contaIsabela.Titular, "-", "ConsultaSaldo atualizado:", contaIsabela.ConsultaSaldo())
	fmt.Println("")

	statusSaque, valor1 := contaHugo.SacarDinheiro(100)
	fmt.Println(contaHugo.Titular, "sacou o valor:", valor1, "-", statusSaque)
	fmt.Println("Conta do:", contaHugo.Titular, "-", "ConsultaSaldo atualizado:", contaHugo.ConsultaSaldo())
	fmt.Println("")

	statusDeposito, valor2 := contaHugo.DepositarDinheiro(300)
	fmt.Println(contaHugo.Titular, "depositou o valor:", valor2, "-", statusDeposito)
	fmt.Println("Conta do:", contaHugo.Titular, "-", "ConsultaSaldo atualizado:", contaHugo.ConsultaSaldo())
	fmt.Println("")

	statusTransf, valorTransf := contaHugo.Transferencia(100, contaIsabela)
	fmt.Println("A transferência ocorreu?", statusTransf)
	fmt.Println("O valor tranferido de", contaHugo.Titular, "para", contaIsabela.Titular, "foi:", valorTransf)
	fmt.Println("Conta do:", contaHugo.Titular, "-", "ConsultaSaldo atualizado:", contaHugo.ConsultaSaldo)
	fmt.Println("Conta da:", contaIsabela.Titular, "-", "ConsultaSaldo atualizado:", contaIsabela.ConsultaSaldo)
	fmt.Println("")


*/

package main

import (
	"fmt"

	"example.com/hello/go/src/GoPOO/contas"
)

func Boleto(conta verificarConta, valorBoleto float64) {
	conta.SacarDinheiro(valorBoleto)
}

type verificarConta interface {
	SacarDinheiro(valor float64) string
}

func main() {

	contaHugo := contas.ContaPoupanca{}
	contaHugo.DepositarDinheiro(100)
	Boleto(&contaHugo, 55) // deve ser utilizado o & na conta para que seja identificada pois não tem uma chamada de identificador na função porque é uma func genérica

	fmt.Println(contaHugo.ConsultaSaldo())

	contaIsabela := contas.ContaCorrente{}
	contaIsabela.DepositarDinheiro(100)
	Boleto(&contaIsabela, 10)

	fmt.Println(contaIsabela.ConsultaSaldo())
}
