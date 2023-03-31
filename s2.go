package main

import "fmt"

func main() {
	var texto string
	fmt.Print("Escreva a palavra ")
	fmt.Scan(&texto)
	s := texto
	fmt.Println("len(s) =", len(s))
}
