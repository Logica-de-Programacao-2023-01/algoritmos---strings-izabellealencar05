package main

import "fmt"

// Solicite ao usuário uma string e imprima somente as suas letras únicas (que aparecem apenas uma vez).
func main() {
	var s string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&s)


	count := make(map[rune]int)
	for _, r := range s {
		count[r]++
	}


	fmt.Print("Letras únicas: ")
	for _, r := range s {
		if count[r] == 1 {
			fmt.Printf("%c ", r)
		}
	}
	fmt.Println()
}
}
