package main

import "fmt"

func main() {
	var x interface{} = 9
	var y interface{} = "Stefania"

	showType(x)
	showType(y)

}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é %t e o valor é %v\n", t, t)

}
