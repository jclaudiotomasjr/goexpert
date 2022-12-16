package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Agencia string `json:"agencia"`
	Numero  string `json:"numero"`
	Saldo   string `json:"saldo"`
}

// func Float64frombytes(bytes []byte) float64 {
// 	bits := binary.LittleEndian.Uint64(bytes)
// 	float := math.Float64frombits(bits)
// 	return float
// }

func main() {
	conta := Conta{Agencia: "8706", Numero: "31894-5", Saldo: "850000"}

	res, err := json.Marshal(conta)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonPuro := []byte(`{"agencia":"8705","numero":"30894-8","saldo":"920000"}`)
	var contaJr Conta
	err = json.Unmarshal(jsonPuro, &contaJr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(contaJr)
}
