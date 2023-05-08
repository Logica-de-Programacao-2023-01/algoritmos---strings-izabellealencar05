package main

import (
	"fmt"
	"unicode"
)

//Solicite ao usuário uma string e informe se ela contém pelo menos um número.
func main() {
	var s string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&s)

	hasNumber := false
	for _, c := range s {
		if unicode.IsDigit(c) {
			hasNumber = true
			break
		}
	}

	if hasNumber {
		fmt.Println("A string contém pelo menos um número.")
	} else {
		fmt.Println("A string não contém nenhum número.")
	}
}
}
