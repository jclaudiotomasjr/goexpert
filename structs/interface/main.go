package main

import "fmt"

type Endereco struct {
	Cep    string
	Rua    string
	Numero string
	Cidade string
	Estado string
}

type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Codigo string
	Nome   string
	Idade  int
	Ativo  bool
}

func (c *Cliente) Desativar() {
	if c.Ativo == true {
		c.Ativo = false
		fmt.Printf("\nCliente %s foi destivado com sucesso!\n", c.Codigo)
		return
	}
	fmt.Printf("\nO clientes já estava desativado")
}

func DesativaCli(p Pessoa) {
	p.Desativar()
}

func main() {

	cl := Cliente{
		Codigo: "J080914",
		Nome:   "Jose Claudio Tomas Junior",
		Idade:  38,
		Ativo:  true,
	}

	fmt.Printf("\nO Cliente %s está ativo %t\n", cl.Codigo, cl.Ativo)
	DesativaCli(&cl)

	fmt.Printf("O cliente está ativo? %t\n", cl.Ativo)
}
