package main

import (
	"fmt"
	"strconv"
)

func main() {

	var x string
	fmt.Print("escreva uma string:")
	fmt.Scan(&x)

	numero, err := strconv.ParseFloat(x, bitSize:64)
	if err != nil {
		fmt.Printf("a string %s não é um ponto flutuante", x)

	}else {
		fmt.Printf("o ponto flutuante é %f", numero)
	}
}
