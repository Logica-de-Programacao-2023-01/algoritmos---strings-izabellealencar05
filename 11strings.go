package main

import (
	"fmt"
	"strings"
)

// Escreva um programa que receba uma string e remova todas as vogais. Informe ao usuário o resultado.func main() {
func main() {
	var s string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&s)

	s = removeVogais(s)

	fmt.Println("String sem vogais:", s)
}

func removeVogais(s string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			return -1
		default:
			return r
		}
	}, s)
}
