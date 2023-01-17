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

	contaDaEvelyn := ContaCorrente{
		titular:       "Evelyn",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.5,
	}

	contaDaLucilene := ContaCorrente{"Lucilene", 222, 111222, 200}
	contaDoGuilherme := ContaCorrente{titular: "Guilherme", saldo: 125.5}

	fmt.Println(contaDaEvelyn)
	fmt.Println(contaDaLucilene)
	fmt.Println(contaDoGuilherme)
}
