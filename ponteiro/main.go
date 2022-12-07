package main

import "fmt"

func main() {
	a := 9
	fmt.Println(a)

	b := &a

	fmt.Println(b)

	*b = 8

	fmt.Println(a)

}
