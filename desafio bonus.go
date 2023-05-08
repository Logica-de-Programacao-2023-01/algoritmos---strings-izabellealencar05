package main

import (
	"fmt"
	"strings"
)

// Escreva um programa que receba uma string e um padrão (outro string) e retorne todos os índices em que o padrão ocorre na string. Informe ao usuário o resultado.
func main() {
	var s, pattern string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&s)

	fmt.Print("Digite um padrão: ")
	fmt.Scanln(&pattern)

	startIndex := 0
	occurrences := []int{}

	for {
		index := strings.Index(s[startIndex:], pattern)
		if index == -1 {
			break
		}
		occurrences = append(occurrences, startIndex+index)
		startIndex += index + len(pattern)
	}

	if len(occurrences) == 0 {
		fmt.Println("O padrão não foi encontrado na string.")
	} else {
		fmt.Println("O padrão foi encontrado nas seguintes posições:", occurrences)
	}
}
}
