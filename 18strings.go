package main

import (
	"fmt"
	"unicode"
)

// Solicite ao usuário uma string e informe se ela contém somente números (0 a 9).
func main() {
	var s string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&s)

	var onlyNumbers = true

	for _, c := range s {
		if !unicode.IsDigit(c) {
			onlyNumbers = false
			break
		}
	}

	if onlyNumbers {
		fmt.Println("A string contém somente números.")
	} else {
		fmt.Println("A string não contém somente números.")
	}
}
}
