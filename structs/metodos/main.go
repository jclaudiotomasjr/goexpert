package main

import "fmt"

type Endereco struct {
	Cep    string
	Rua    string
	Numero int
	Bairo  string
	Cidade string
	Estado string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func (c Cliente) Desativar() {
	if c.Ativo == true {
		c.Ativo = false
		fmt.Printf("O cliente %s, foi desativado com sucesso!\n", c.Nome)
		return
	}
	fmt.Printf("o Cliente %s, já está desativado!\n", c.Nome)
}

func main() {

	cl := Cliente{
		Nome:  "Jose Claudio Tomas Jr",
		Idade: 38,
		Ativo: false,
	}

	fmt.Printf("\nNome: %s\nIdade: %d\nAtivo: %t\n\n", cl.Nome, cl.Idade, cl.Ativo)

	cl.Desativar()

}
