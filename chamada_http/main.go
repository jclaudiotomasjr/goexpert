package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	req, err := http.Get("https://www.alura.com.br/")
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

}
