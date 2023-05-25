package main

import (
	"fmt"
	"strings"
)

func main() {
	// Solicitar duas strings ao usuário
	fmt.Print("Digite a primeira string: ")
	var str1 string
	fmt.Scanln(&str1)

	fmt.Print("Digite a segunda string: ")
	var str2 string
	fmt.Scanln(&str2)

	// Verificar se a segunda string é uma substring da primeira
	if ehSubstring(str1, str2) {
		fmt.Println("A segunda string é uma substring da primeira!")
	} else {
		fmt.Println("A segunda string não é uma substring da primeira.")
	}
}

// Função para verificar se a segunda string é uma substring da primeira
func ehSubstring(str1, str2 string) bool {
	return strings.Contains(str1, str2)
}
