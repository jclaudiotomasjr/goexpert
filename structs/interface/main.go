package main

import "fmt"

type Endereco struct {
	Cep    string
	Rua    string
	Numero string
	Cidade string
	Estado string
}

type Cliente struct {
	Codigo string
	Nome   string
	Idade  int
	Ativo  bool
}

func (c Cliente) Desativar() {
	if c.Ativo == true {
		c.Ativo = false
		fmt.Printf("Cliente %s\n destivado com sucesso!\n", c.Nome)
		return
	}
	fmt.Printf("O clientes já estava desativado")
}

func main() {

}
