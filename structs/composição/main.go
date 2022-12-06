package main

import "fmt"

type Endereco struct {
	Rua    string
	Numero int
	Bairro string
	Cidade string
	Estado string
}
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {

	cl := Cliente{
		Nome:  "Jose Claudio Tomás Jr",
		Idade: 39,
		Ativo: true,
	}

	cl.Rua = "Barão de Santo Angelo"
	cl.Numero = 690
	cl.Bairro = "Engenho de Dentro"
	cl.Cidade = "Rio de Janeiro"
	cl.Estado = "Rio de Janeiro"
	fmt.Printf("\nNome: %s\nIdade: %d\nAtivo: %t\nRua: %s %d\nBairo: %s\nCidade: %s\nEstado: %s\n", cl.Nome, cl.Idade, cl.Ativo, cl.Rua, cl.Numero, cl.Bairro, cl.Cidade, cl.Estado)

}
