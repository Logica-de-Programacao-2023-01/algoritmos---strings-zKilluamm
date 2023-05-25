package main

import (
	"fmt"
)

func main() {
	// Solicitar uma string ao usuário
	fmt.Print("Digite uma string: ")
	var entrada string
	fmt.Scanln(&entrada)

	// Inverter a string
	invertida := inverterString(entrada)

	// Imprimir a string invertida
	fmt.Println("String invertida:", invertida)
}

// Função para inverter uma string
func inverterString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
