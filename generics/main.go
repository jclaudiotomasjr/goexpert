package main

import "fmt"

func soma[T int | int32 | int64 | float32 | float64](m map[string]T) T {
	var total T
	for _, v := range m {
		total += v
	}
	return total
}

func main() {

	cl0 := map[string]int{"Claudio": 15000, "Stefania": 15000, "Vitor": 10000, "Yan": 10000}
	cl1 := map[string]int32{"Claudio": 15000, "Stefania": 15000, "Vitor": 10000, "Yan": 10000}
	cl2 := map[string]int64{"Claudio": 15000, "Stefania": 15000, "Vitor": 10000, "Yan": 10000}
	cl3 := map[string]float32{"Claudio": 15000.50, "Stefania": 15000.50, "Vitor": 10000, "Yan": 10000}
	cl4 := map[string]float64{"Claudio": 15000.50, "Stefania": 15000.50, "Vitor": 10000, "Yan": 10000}
	fmt.Println(soma(cl0))
	fmt.Println(soma(cl1))
	fmt.Println(soma(cl2))
	fmt.Println(soma(cl3))
	fmt.Println(soma(cl4))

}
