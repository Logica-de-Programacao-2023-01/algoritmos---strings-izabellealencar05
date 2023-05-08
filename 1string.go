package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&s)

	s = strings.ToUpper(s)

	fmt.Println("String convertida:", s)
}
