package main

import (
	"fmt"
	"unicode"
)

func main() {
	// Solicitar uma string ao usuário
	fmt.Print("Digite uma string: ")
	var entrada string
	fmt.Scanln(&entrada)

	// Verificar se a string está em CamelCase
	if ehCamelCase(entrada) {
		fmt.Println("A string está em CamelCase!")
	} else {
		fmt.Println("A string não está em CamelCase.")
	}

	// Contar a quantidade de palavras na string
	quantidadePalavras := contarPalavrasCamelCase(entrada)
	fmt.Println("Quantidade de palavras:", quantidadePalavras)
}

// Função para verificar se uma string está em CamelCase
func ehCamelCase(s string) bool {
	// Verificar se a primeira letra é maiúscula
	if len(s) > 0 && !unicode.IsUpper(rune(s[0])) {
		return false
	}

	// Verificar se há letras minúsculas seguidas de maiúsculas
	for i := 1; i < len(s); i++ {
		if unicode.IsLower(rune(s[i-1])) && unicode.IsUpper(rune(s[i])) {
			return false
		}
	}

	return true
}

// Função para contar a quantidade de palavras em CamelCase
func contarPalavrasCamelCase(s string) int {
	quantidade := 1
	for i := 1; i < len(s); i++ {
		if unicode.IsUpper(rune(s[i])) {
			quantidade++
		}
	}
	return quantidade
}
