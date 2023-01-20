package main

import (
	"fmt"

	"github.com/alura/banco/contas"
)

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDaPati := contas.ContaCorrente{}

	fmt.Println(contaDoDenis)
	fmt.Println(contaDaPati)

	contaDoDenis.Depositar(100)
	fmt.Println(contaDoDenis.ObterSaldo())

	contaDoDenis.Sacar(55)
	fmt.Println(contaDoDenis.ObterSaldo())

	contaDoDenis.Sacar(1000)
	fmt.Println(contaDoDenis.ObterSaldo())
}
