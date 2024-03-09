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

}
