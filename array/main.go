package main

import "fmt"

func main() {

	var meuArray [3]int

	meuArray[0] = 9
	meuArray[1] = 14
	meuArray[2] = 8

	for i, valor := range meuArray {
		fmt.Printf("O valor do indice é %d e o valor é %d\n", i, valor)
	}
	fmt.Println(meuArray[len(meuArray)-1])

}
