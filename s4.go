package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra string
	fmt.Print("Digita um palavra: ")
	fmt.Scan(&palavra)
	fmt.Println(strings.ToUpper(palavra))
}
