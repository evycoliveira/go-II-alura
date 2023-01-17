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

	contaDaEvelyn2 := ContaCorrente{
		titular:       "Evelyn",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.5,
	}

	// Segunda forma de atribuição de todos os valores
	contaDaLucilene := ContaCorrente{"Lucilene", 222, 111222, 200}
	contaDaLucilene2 := ContaCorrente{"Lucilene", 222, 111222, 200}
	// Terceira forma com passagem parcial de valores
	contaDoGuilherme := ContaCorrente{titular: "Guilherme", saldo: 125.5}

	// True
	fmt.Println(contaDaEvelyn == contaDaEvelyn2)
	// True
	fmt.Println(contaDaLucilene == contaDaLucilene2)
	fmt.Println(contaDaEvelyn)
	fmt.Println(contaDaLucilene)
	fmt.Println(contaDoGuilherme)

	// É necessaria tornar um ponteiro para que na linha do new, a referência não tenha erro
	var contaDaCristiane *ContaCorrente
	contaDaCristiane = new(ContaCorrente)
	contaDaCristiane.titular = "Cristiane"
	contaDaCristiane.saldo = 500

	var contaDaCristiane2 *ContaCorrente
	contaDaCristiane2 = new(ContaCorrente)
	contaDaCristiane.titular = "Cristiane"
	contaDaCristiane.saldo = 500

	// Resultado será o endereço de memoria e os valores atribuídos
	fmt.Println(contaDaCristiane)
	// Resultado será o endereço os valores atribuídos
	fmt.Println(*contaDaCristiane)
	// False: mesmo sendo os mesmos valores, por conta dos diferentes endereços de memoria
	fmt.Println(contaDaCristiane == contaDaCristiane2)
	// Endereços de memória diferentes
	fmt.Println(*contaDaCristiane == *contaDaCristiane2)
}
