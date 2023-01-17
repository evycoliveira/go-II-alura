package main

import "fmt"

func main() {
	type ContaCorrente struct {
		titular       string
		numeroAgencia int
		numeroConta   int
		saldo         float64
	}

	fmt.Println(ContaCorrente{})

	// Primeira forma de atribuição de valores
	contaDaEvelyn := ContaCorrente{
		titular:       "Evelyn",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.5,
	}

	// Segunda forma de atribuição de todos os valores
	contaDaLucilene := ContaCorrente{"Lucilene", 222, 111222, 200}
	// Terceira forma com passagem parcial de valores
	contaDoGuilherme := ContaCorrente{titular: "Guilherme", saldo: 125.5}

	fmt.Println(contaDaEvelyn)
	fmt.Println(contaDaLucilene)
	fmt.Println(contaDoGuilherme)

	// É necessaria tornar um ponteiro para que na linha do new, a referência não tenha erro
	var contaDaCristiane *ContaCorrente
	contaDaCristiane = new(ContaCorrente)
	contaDaCristiane.titular = "Cristiane"
	contaDaCristiane.saldo = 500

	// Resultado será o endereço de memoria e os valores atribuídos
	fmt.Println(contaDaCristiane)
	// Resultado será o endereço os valores atribuídos
	fmt.Println(*contaDaCristiane)
}
