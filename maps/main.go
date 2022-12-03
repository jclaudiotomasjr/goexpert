package main

import "fmt"

func main() {
	salarios := map[string]int{"Vitor": 12000, "Junior": 18000, "Stefania": 14000, "Yan": 12000, "Jack": 500}

	fmt.Println(salarios["Vitor"])

	for nome, salario := range salarios {
		fmt.Printf("Seu nome é %v e seu Salario é %d\n", nome, salario)
	}
}
