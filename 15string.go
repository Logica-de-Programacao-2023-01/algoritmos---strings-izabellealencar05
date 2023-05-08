package main

import (
	"fmt"
	"strings"
)

// Solicite ao usuário uma string e substitua todas as vogais por '*' (asterisco).
func main() {
	var s string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&s)

	s = strings.Map(func(r rune) rune {
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			return '*'
		default:
			return r
		}
	}, s)

	fmt.Println("String com as vogais substituídas:", s)
}
}
