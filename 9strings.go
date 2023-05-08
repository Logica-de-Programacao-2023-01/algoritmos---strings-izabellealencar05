package main

import (
	"fmt"
	"strings"
)

// Solicite ao usuário uma string e substitua todas as ocorrências de uma letra por outra informada pelo usuário.
func main() {
	var s, oldChar, newChar string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&s)

	fmt.Print("Digite a letra que deseja substituir: ")
	fmt.Scanln(&oldChar)

	fmt.Print("Digite a letra pela qual deseja substituir: ")
	fmt.Scanln(&newChar)

	s = strings.ReplaceAll(s, oldChar, newChar)

	fmt.Println("Nova string:", s)
}
}
