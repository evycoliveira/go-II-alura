package main

import (
	"fmt"

	"github.com/alura/banco/clientes"
	"github.com/alura/banco/contas"
)

func main() {
	contaDoBruno := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Bruno",
		CPF:       "123.222.444-77",
		Profissao: "Desenvolvedor",
	},
		NumeroAgencia: 123,
		NumeroConta:   123456,
		Saldo:         100,
	}

	fmt.Println(contaDoBruno)
}
