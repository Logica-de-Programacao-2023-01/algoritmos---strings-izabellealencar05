package main

import "fmt"

// Escreva um programa que receba uma string e inverta a ordem dos caracteres. Informe ao usuário o resultado.
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var s string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&s)
	reversed := reverseString(s)
	fmt.Println("String invertida:", reversed)
}
