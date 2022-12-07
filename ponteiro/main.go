package main

import "fmt"

func soma(a, b *int) int {
	*a = 14
	*b = 3
	return *a + *b
}

func main() {
	a, b := 8, 9

	fmt.Println(soma(&a, &b))
	println(a, b)

}
