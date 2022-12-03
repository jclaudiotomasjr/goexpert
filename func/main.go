package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := soma(14, 39)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(valor)

}

func soma(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma dos valores é maiore do que 50")
	}
	return a + b, nil
}
