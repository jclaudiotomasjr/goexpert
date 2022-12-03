package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Cargo string
	Ativo bool
}

func main() {

	Claudio := Cliente{
		Nome:  "Claudio Tomás Jr",
		Idade: 38,
		Cargo: "Programador",
		Ativo: true,
	}

	fmt.Printf("Nome: %s\nIdade: %d\nCargo: %s\nAtivo: %t\n", Claudio.Nome, Claudio.Idade, Claudio.Cargo, Claudio.Ativo)

}
