package main

import (
	"fmt"
	"strings"
)

func main() {

	var palavra, caractere string

	fmt.Print("Informe uma palavra: ")
	fmt.Scan(&palavra)
	fmt.Print("Informe um caractere: ")
	fmt.Scan(&caractere)

	count := strings.Count(palavra, caractere)
	fmt.Printf("O caractere '%s' aparece %d vezes na string '%s' .\n", caractere, count, palavra)
}
