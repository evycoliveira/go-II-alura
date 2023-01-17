package main

import (
	"fmt"

	c "github.com/alura/banco/contas"
)

func main() {

	fmt.Println(c.ContaCorrente{})
	contaDaSilvia := c.ContaCorrente{}
	contaDaSilvia.Titular = "Silvia"
	contaDaSilvia.Saldo = 500

	fmt.Println(contaDaSilvia.Saldo)
	fmt.Println(contaDaSilvia.Sacar(200))
	fmt.Println(contaDaSilvia.Saldo)

	statusDeposito, valor := contaDaSilvia.Depositar(2000)
	fmt.Println(statusDeposito, valor)

	contaDaSilvia = c.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := c.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	statusTransferenciaSilvia := contaDaSilvia.Transferir(200, &contaDoGustavo)
	statusTransferenciaGustavo := contaDoGustavo.Transferir(-200, &contaDaSilvia)

	fmt.Println(statusTransferenciaSilvia)
	fmt.Println(statusTransferenciaGustavo)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
}
