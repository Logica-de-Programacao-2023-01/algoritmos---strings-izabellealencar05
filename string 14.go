package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var x string
	fmt.Print("escreva uma string:")
	fmt.Scan(&x)

	_, err := strconv.Atoi(x)
	if err != nil {
		fmt.Printf("a string %s nao é uma sequencia numerica", x)

	} else {
		numeros := strings.Split(x, "")
		decrescente := true
		for i := 1; i < len(numeros); i++ {
			anterior := numeros[i-1]
			atual := numeros[i]
			if anterior <= atual {
				decrescente = false
				break
			}
		}
	}
}
