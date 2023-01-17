package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	fmt.Println(ContaCorrente{})
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.saldo)
	fmt.Println(contaDaSilvia.Sacar(200))
	fmt.Println(contaDaSilvia.saldo)

	statusDeposito, valor := contaDaSilvia.Depositar(2000)
	fmt.Println(statusDeposito, valor)

	contaDaSilvia = ContaCorrente{titular: "Silvia", saldo: 300}
	contaDoGustavo := ContaCorrente{titular: "Gustavo", saldo: 100}

	statusTransferenciaSilvia := contaDaSilvia.Transferir(200, &contaDoGustavo)
	statusTransferenciaGustavo := contaDoGustavo.Transferir(-200, &contaDaSilvia)

	fmt.Println(statusTransferenciaSilvia)
	fmt.Println(statusTransferenciaGustavo)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente!"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso!", c.saldo
	} else {
		return "Valor do depósito menor que zero!", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}
