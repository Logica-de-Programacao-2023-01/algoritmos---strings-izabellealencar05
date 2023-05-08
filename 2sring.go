package main

import (
	"fmt"
	"strings"
)

//Escreva um programa que receba uma string e remova todos os espaços em branco. Informe ao usuário o resultado.

func main() {
	var s string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&s)

	s = strings.ReplaceAll(s, " ", "")

	fmt.Println("String sem espaços:", s)
}
