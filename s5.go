package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	var n int

	fmt.Print("Digite uma string: ")
	fmt.Scan(&str)

	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	if n > len(str) {
		fmt.Println("O número é maior que o tamanho da string.")
		return
	}

	newStr := strings.ToUpper(str[:n]) + str[n:]

	fmt.Println("Nova string:", newStr)
}
