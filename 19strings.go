package main

import "fmt"

// Solicite ao usuário uma string e imprima a mesma string invertida.
func main() {
	var s string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&s)


	r := []rune(s)


	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	 
	s = string(r)

	fmt.Println("String invertida:", s)
}
}
