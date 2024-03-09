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


*/

package main

import "fmt"

//criação de uma estrutura com variáveis que serão usadas com frequencia para não serem repetidas
type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	contaHugo := ContaCorrente{titular: "Hugo", numeroAgencia: 555, numeroConta: 1234, saldo: 150.25} // forma extensa de passar os dados para a estrutura
	contaIsabela := ContaCorrente{"Isabela", 500, 4321, 100.00}                                       // forma curta de passar dados para a estrutura

	fmt.Println(contaHugo)
	fmt.Println(contaIsabela)

	//Criando uma nova conta/variável na estrutura
	// sem colocar * o Go não entende de onde puxar a informação; se é da estrutura ou da nova ContaCorrente criada depois
	var contaLaura *ContaCorrente
	// criação de um novo bloco de dados para a estrutura
	contaLaura = new(ContaCorrente)
	// atribuição dos dados exigidos pela estrutura para a Nova ContaCorrente
	contaLaura.titular = "Laura"
	contaLaura.numeroAgencia = 2222

	// irá gerar no console a informação sem rotulação porém com os dados informados
	fmt.Println(contaLaura)
	// irá gerar no console a informação com rotulação (nome atribuido) com os dados informados
	fmt.Println(*contaLaura)

}
