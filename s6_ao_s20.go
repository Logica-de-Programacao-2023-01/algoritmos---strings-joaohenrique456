package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// q.6

func main() {
	var frase string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&frase)

	numpalavras := Contadorpalavras(frase)

	fmt.Println("A string contém", numpalavras, "palavras.")
}

func Contadorpalavras(frase string) int {
	Espaço := strings.TrimSpace(frase)
	palavras := strings.Fields(Espaço)
	return len(palavras)
}

// q.7

func main() {

	var texto string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&texto)

	num := false

	for _, Letra := range texto {
		if unicode.IsDigit(Letra) {
			num = true
			break
		}
	}

	if num {
		fmt.Println("A string contém pelo menos um número.")
	} else {
		fmt.Println("A string não contém números.")
	}
}

// q.8

func main() {

	var texto string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&texto)

	stringinvertida := stringinvertida(texto)

	fmt.Println("String invertida:", stringinvertida)
}

func stringinvertida(input string) string {
	escrita := []rune(input)
	tamanho := len(escrita)

	for i := 0; i < tamanho/2; i++ {
		escrita[i], escrita[tamanho-1-i] = escrita[tamanho-1-i], escrita[i]
	}

	return string(escrita)
}

// q.9

func main() {

	var nome, Al, Nl string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&nome)

	fmt.Print("Digite o caractere a ser substituído: ")
	fmt.Scanln(&Al)

	fmt.Print("Digite o novo caractere: ")
	fmt.Scanln(&Nl)

	NS := strings.ReplaceAll(nome, Al, Nl)

	fmt.Println("String atualizada:", NS)
}

// q.10

func main() {

	var str1, str2 string

	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&str1)

	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&str2)

	if éanagra(str1, str2) {
		fmt.Println("As strings são anagramas.")
	} else {
		fmt.Println("As strings não são anagramas.")
	}
}

func éanagra(str1, str2 string) bool {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	str1 = organizaçãodestring(str1)
	str2 = organizaçãodestring(str2)

	return str1 == str2
}

func organizaçãodestring(str string) string {
	letras := strings.Split(str, "")
	sort.Strings(letras)
	return strings.Join(letras, "")
}

// q.11

func main() {

	var nomes string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&nomes)

	vogaisremovidas := vogaisremovidas(nomes)

	fmt.Println("String sem vogais:", vogaisremovidas)
}

func vogaisremovidas(input string) string {
	vogais := "aeiouAEIOU"
	resultado := ""

	for _, letra := range input {
		if !strings.ContainsRune(vogais, letra) {
			resultado += string(letra)
		}
	}
	return resultado
}

// q.12

func main() {

	var nomes string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&nomes)

	if palindromo(nomes) {
		fmt.Println("A string é um palíndromo.")
	} else {
		fmt.Println("A string não é um palíndromo.")
	}
}

func palindromo(nome string) bool {
	nome = strings.ToLower(strings.ReplaceAll(nome, " ", ""))

	tamanho := len(nome)
	for i := 0; i < tamanho/2; i++ {
		if nome[i] != nome[tamanho-1-i] {
			return false
		}
	}

	return true
}

// q.13

func main() {

	var nomes string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&nomes)

	if sequencia(nomes) {
		fmt.Println("A string é uma sequência numérica crescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica crescente.")
	}
}

func sequencia(input string) bool {
	tamanho := len(input)
	if tamanho < 2 {
		return false
	}

	for i := 1; i < tamanho; i++ {
		atual, err1 := strconv.Atoi(string(input[i]))
		anterior, err2 := strconv.Atoi(string(input[i-1]))

		if err1 != nil || err2 != nil || atual <= anterior {
			return false
		}
	}

	return true
}

// q.14

func main() {

	var nomes string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&nomes)

	if NS(nomes) {
		fmt.Println("A string é uma sequência numérica decrescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica decrescente.")
	}
}
func NS(nomes string) bool {
	tamanho := len(nomes)
	if tamanho < 2 {
		return false
	}
	for i := 1; i < tamanho; i++ {
		atual, err1 := strconv.Atoi(string(nomes[i]))
		anterior, err2 := strconv.Atoi(string(nomes[i-1]))

		if err1 != nil || err2 != nil || atual >= anterior {
			return false
		}
	}
	return true
}

// q.15

func main() {

	var nomes string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&nomes)

	Nstring := VR(nomes)

	fmt.Println("String com vogais substituídas:", Nstring)
}

func VR(input string) string {
	vogais := "aeiouAEIOU"
	Repositor := strings.NewReplacer(vogais, "*")

	return Repositor.Replace(input)
}

// q.16

func main() {

	var nomes string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&nomes)

	Nstring := VR(nomes)

	fmt.Println("String com vogais substituídas:", Nstring)
}

func VR(input string) string {
	vogais := "aeiouAEIOU"
	Repositor := strings.NewReplacer(vogais, "*")

	return Repositor.Replace(input)
}

// q.17

func main() {

	var nomes string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&nomes)

	LU := getUniqueLetters(nomes)

	fmt.Println("Letras únicas na string:", LU)
}

func getUniqueLetters(input string) string {
	unicas := ""
	javistas := make(map[rune]int)

	for _, letras := range input {
		javistas[letras]++
	}

	for _, char := range input {
		if javistas[char] == 1 {
			unicas += string(char)
		}
	}

	return unicas
}

// q.18

func main() {

	var nomes string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&nomes)

	if CN(nomes) {
		fmt.Println("A string contém somente números.")
	} else {
		fmt.Println("A string não contém somente números.")
	}
}

func CN(nomes string) bool {
	_, err := strconv.Atoi(nomes)
	return err == nil
}

// q.19

func main() {

	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	reversedString := reverseString(input)

	fmt.Println("String invertida:", reversedString)
}

func reverseString(input string) string {
	runes := []rune(input)
	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// q.20

func main() {

	var nomes string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&nomes)

	isCamelCase := checkCamelCase(nomes)
	numplavras := countWords(nomes)

	if isCamelCase {
		fmt.Println("A string está em camelCase.")
	} else {
		fmt.Println("A string não está em camelCase.")
	}

	fmt.Println("Número de palavras:", numplavras)
}

func checkCamelCase(input string) bool {
	if len(input) == 0 {
		return false
	}

	firstChar := input[0]
	if !maiuscula(firstChar) {
		return false
	}

	for i := 1; i < len(input); i++ {
		char := input[i]
		if maiuscula(char) {
			return false
		}
	}

	return true
}

func maiuscula(char byte) bool {
	return char >= 'A' && char <= 'Z'
}

func countWords(input string) int {
	words := strings.FieldsFunc(input, func(r rune) bool {
		return !maiuscula(byte(r))
	})
	return len(words)
}