package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type BuscaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

//URL: http://viacep.com.br/ws/01001000/json/

func main() {

	for _, cep := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
			return
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
			return
		}

		var data BuscaCep

		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
			return
		}
		fmt.Println(data)

		file, err := os.Create("cep.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %s", err)
			return
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s,\nLocalidade: %s, \nUF: %s, \n ", data.Cep, data.Localidade, data.Uf))

	}

}
