package main

import (
	"fmt"
	"strings"
)

func main() {
	// Solicitar uma string ao usuário
	fmt.Print("Digite uma palavra ou frase: ")
	var entrada string
	fmt.Scanln(&entrada)

	// Remover espaços em branco e converter para minúsculas
	entrada = strings.ToLower(strings.ReplaceAll(entrada, " ", ""))

	// Verificar se a string é um palíndromo
	if ehPalindromo(entrada) {
		fmt.Println("A palavra/frase é um palíndromo!")
	} else {
		fmt.Println("A palavra/frase não é um palíndromo.")
	}
}

// Função para verificar se uma string é um palíndromo
func ehPalindromo(s string) bool {
	// Percorrer a string da esquerda para a direita
	// e da direita para a esquerda, comparando os caracteres
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
