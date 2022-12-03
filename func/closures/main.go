package main

import "fmt"

func main() {

	total := func() int {
		return sumAll(1, 2, 3, 4, 8, 9, 11, 14, 20) * 2
	}()
	fmt.Println("A soma total dos valores é: ", total)
}

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total

}
