package main

import "fmt"

func main() {

	fmt.Println(sumAll(2, 3, 8, 9, 14, 1, 4))

}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
