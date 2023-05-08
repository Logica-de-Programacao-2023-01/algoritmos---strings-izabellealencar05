package main

import (
	"fmt"
	"strconv"
)

// Solicite ao usuário uma string e informe se ela é uma sequência numérica crescente (exemplo: "123" ou "4567").
func main() {
	var s string

	fmt.Print("Digite uma sequência numérica: ")
	fmt.Scanln(&s)

	for i := 1; i < len(s); i++ {
		prev, _ := strconv.Atoi(string(s[i-1]))
		curr, _ := strconv.Atoi(string(s[i]))
		if curr != prev+1 {
			fmt.Println("Não é uma sequência numérica crescente.")
			return
		}
	}

	fmt.Println("É uma sequência numérica crescente.")
}
}
