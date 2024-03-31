/*
PROJETO CURSO ALURA - ORIENTAÇÂO A OBJETOS

Criação de um gerenciador de contas de um banco

## tipo de atribuição de variável simples:

var Titular string = "Hugo"
var numeroAgencia int = 555
var numeroConta int = 123456
var saldo float64 = 150.25

ou

## forma reduzida de declaração de variável onde o Go reconhece o tipo sem precisar explicitar

Titular := "Hugo"
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
	contaLaura.Titular = "Laura"
	contaLaura.numeroAgencia = 222

// irá gerar no console a informação SEM rotulação porém com os dados informados
	fmt.Println(contaLaura)
	// console // &{Laura 222 0 0} => onde o & é um endereço

// irá gerar no console a informação COM rotulação (nome atribuido) com os dados informados
	fmt.Println(*contaLaura)
	// console // {Laura 222 0 0}


# metodo de comparação e verificação se os 2 blocos de dados (conteúdo) são iguais ou diferentes

	contaHugo := ContaCorrente{Titular: "Hugo", numeroAgencia: 555, numeroConta: 1234, saldo: 150.25}
	contaHugo2 := ContaCorrente{Titular: "Hugo", numeroAgencia: 555, numeroConta: 1234, saldo: 150.25}

	fmt.Println(contaHugo == contaHugo2)
	// console // true (retorna um boolenado informando que é verdandeiro apesar de serem contas diferentes pois possuem os mesmos dados)



	contaHugo := ContaCorrente{Titular: "Hugo", numeroAgencia: 555, numeroConta: 1234, saldo: 150.25}
	fmt.PrintLn(contaHugo.saldo)
	// console // 150.25

	valorSaque = 50
	contaHugo.saldo = contaHugo.saldo - valorSaque
	fmt.Println(contaHugo.saldo)
	// console // 100.25


	[Info  - 12:28:25] 2024/03/24 12:28:25 go/packages.Load #188
	snapshot=1487
	directory=file:///C:/Users/Hugo/go/src/GoPOO

	package="example.com/hello/go/src/GoPOO"

	files=[C:\Users\Hugo\go\src\GoPOO\main.go]


*/

package main

import (
	"fmt"

	"example.com/hello/go/src/GoPOO/clientes"
	"example.com/hello/go/src/GoPOO/contas"
)

func main() {

	//criando o campo de cliente e conta baseada na composição das 2 structs importadas
	clienteHugo := clientes.Titular{"Hugo", "123.123.123-12"}
	contaHugo := contas.ContaCorrente{clienteHugo, 555, 1234, 500}

	clienteIsabela := clientes.Titular{"Isabela", "111.222.333.44"}
	contaIsabela := contas.ContaCorrente{clienteIsabela, 500, 4321, 200}

	// contaHugo := contas.ContaCorrente{Titular: "Hugo", NumeroAgencia: 555, NumeroConta: 1234, Saldo: 500}
	// contaIsabela := contas.ContaCorrente{Titular: "Isabela", NumeroAgencia: 500, NumeroConta: 4321, Saldo: 100}

	fmt.Println("Dados das Contas:")
	fmt.Println(contaHugo)
	fmt.Println(contaIsabela)
	fmt.Println("")

	fmt.Println("Conta do:", contaHugo.Titular, "-", "Saldo atualizado:", contaHugo.Saldo)
	fmt.Println("Conta da:", contaIsabela.Titular, "-", "Saldo atualizado:", contaIsabela.Saldo)
	fmt.Println("")

	statusSaque, valor2 := contaHugo.SacarDinheiro(100)
	fmt.Println(contaHugo.Titular, "sacou o valor:", valor2, "-", statusSaque)
	fmt.Println("Conta do:", contaHugo.Titular, "-", "Saldo atualizado:", contaHugo.Saldo)
	fmt.Println("")

	statusDeposito, valor := contaHugo.DepositarDinheiro(300)
	fmt.Println(contaHugo.Titular, "depositou o valor:", valor, "-", statusDeposito)
	fmt.Println("Conta do:", contaHugo.Titular, "-", "Saldo atualizado:", contaHugo.Saldo)
	fmt.Println("")

	statusTransf, valorTransf := contaHugo.Transferencia(100, &contaIsabela)
	fmt.Println("A transferência ocorreu?", statusTransf)
	fmt.Println("O valor tranferido de", contaHugo.Titular, "para", contaIsabela.Titular, "foi:", valorTransf)
	fmt.Println("Conta do:", contaHugo.Titular, "-", "Saldo atualizado:", contaHugo.Saldo)
	fmt.Println("Conta da:", contaIsabela.Titular, "-", "Saldo atualizado:", contaIsabela.Saldo)
	fmt.Println("")
}
