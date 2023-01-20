package main

import (
	"fmt"

	"github.com/alura/banco/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDaPati := contas.ContaCorrente{}

	fmt.Println(contaDoDenis)
	fmt.Println(contaDaPati)

	contaDaPati.Depositar(3000)
	fmt.Println(contaDoDenis.ObterSaldo())

	contaDoDenis.Depositar(10000)
	fmt.Println(contaDoDenis.ObterSaldo())

	contaDoDenis.Sacar(55)
	fmt.Println(contaDoDenis.ObterSaldo())

	contaDoDenis.Sacar(1000)
	fmt.Println(contaDoDenis.ObterSaldo())

	PagarBoleto(&contaDoDenis, 60)
	fmt.Println(contaDoDenis.ObterSaldo())

	PagarBoleto(&contaDaPati, 100)
	fmt.Println(contaDaPati.ObterSaldo())
}
