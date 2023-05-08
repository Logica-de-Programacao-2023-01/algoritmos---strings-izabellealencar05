package main

import (
	"fmt"
	"strings"
)

// Escreva um programa que receba uma string e conte quantas palavras ela contém. Informe ao usuário o resultado.
func main() {
	var s string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&s)

	// remove espaços extras no início e fim da string
	s = strings.TrimSpace(s)

	// conta o número de palavras
	numPalavras := 1
	for _, c := range s {
		if c == ' ' {
			numPalavras++
		}
	}

	fmt.Printf("A string contém %d palavras.\n", numPalavras)
}

}
