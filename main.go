package main

import (
	"fmt"

	"github.com/alura/banco/clientes"
	"github.com/alura/banco/contas"
)

func main() {
	clienteBruno := clientes.Titular{"Bruno", "123.222.333-88", "Desenvolvedor Go"}

	contaDoBruno := contas.ContaCorrente{clienteBruno, 123, 123456, 100}
	fmt.Println(contaDoBruno)
}
