package main

import "fmt"

//Escreva um programa que receba uma string e substitua todas as ocorrências desse caractere na string por outro caractere especificado pelo usuário. Informe ao usuário o resultado.

func replaceChar(s string, oldChar rune, newChar rune) string {
	r := []rune(s)
	for i, c := range r {
		if c == oldChar {
			r[i] = newChar
		}
	}
	return string(r)
}

func main() {
	var s string
	var oldChar, newChar rune

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&s)

	fmt.Print("Digite o caractere a ser substituído: ")
	fmt.Scanf("%c", &oldChar)

	fmt.Print("Digite o novo caractere: ")
	fmt.Scanf("%c", &newChar)

	s = replaceChar(s, oldChar, newChar)

	fmt.Println("String resultante:", s)
}
