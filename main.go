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
}
