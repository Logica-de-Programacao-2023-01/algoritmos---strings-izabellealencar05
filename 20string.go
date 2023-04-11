package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Print("escreva uma string:")
	fmt.Scan(&x)

	if string(x[0]) == strings.ToLower(string(x[0])) {
		contador := 1
		for _, letra := range x {
			if string(letra) == strings.ToLower(string(letra)) {
				contador++
			}
		}
		fmt.Printf("a string %s possui %dpalavras", x, contador)
	} else {
		fmt.Println("nao esta em camelcase")

	}
}
