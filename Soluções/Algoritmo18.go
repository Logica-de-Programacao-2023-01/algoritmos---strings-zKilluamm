package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Solicitar uma string ao usuário
	fmt.Print("Digite uma string: ")
	var entrada string
	fmt.Scanln(&entrada)

	// Verificar se a string contém somente números
	if contemSomenteNumeros(entrada) {
		fmt.Println("A string contém somente números!")
	} else {
		fmt.Println("A string não contém somente números.")
	}
}

// Função para verificar se uma string contém somente números (0 a 9)
func contemSomenteNumeros(s string) bool {
	pattern := "^[0-9]+$"
	match, _ := regexp.MatchString(pattern, s)
	return match
}
