package main

import (
	"fmt"
	"strings"
)

// Solicite ao usuário duas strings e informe se a segunda é uma substring da primeira
func main() {
	var s1, s2 string

	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&s1)

	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&s2)

	if strings.Contains(s1, s2) {
		fmt.Println("A segunda string é uma substring da primeira.")
	} else {
		fmt.Println("A segunda string não é uma substring da primeira.")
	}
}
