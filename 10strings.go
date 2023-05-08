package main

import (
	"fmt"
	"sort"
	"strings"
)

// Escreva um programa que receba duas strings e verifique se elas são anagramas. Informe ao usuário se são ou não.
func main() {
	var str1, str2 string

	fmt.Println("Digite a primeira string:")
	fmt.Scanln(&str1)
	fmt.Println("Digite a segunda string:")
	fmt.Scanln(&str2)

	str1 = strings.ToLower(strings.ReplaceAll(str1, " ", ""))
	str2 = strings.ToLower(strings.ReplaceAll(str2, " ", ""))

	slice1 := strings.Split(str1, "")
	slice2 := strings.Split(str2, "")

	sort.Strings(slice1)
	sort.Strings(slice2)

	if len(slice1) != len(slice2) {
		fmt.Println("Não são anagramas")
		return
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			fmt.Println("Não são anagramas")
			return
		}
	}

	fmt.Println("São anagramas")
}
