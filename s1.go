package main

import "fmt"

func main() {

	var s1, s2 string

	fmt.Print("Informe a primeira string: ")
	fmt.Scan(&s1)
	fmt.Print("Informe a segunda string: ")
	fmt.Scan(&s2)

	s3 := s1 + " " + s2
	fmt.Println(s3)
}
