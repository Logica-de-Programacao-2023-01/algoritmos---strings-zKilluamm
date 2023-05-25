package main

import (
	"fmt"
)

func main() {
	// Solicitar uma string ao usuário
	fmt.Print("Digite uma sequência numérica: ")
	var entrada string
	fmt.Scanln(&entrada)

	// Verificar se a string é uma sequência numérica crescente
	if ehSequenciaCrescente(entrada) {
		fmt.Println("A sequência é numérica crescente!")
	} else {
		fmt.Println("A sequência não é numérica crescente.")
	}
}

// Função para verificar se uma string é uma sequência numérica crescente
func ehSequenciaCrescente(s string) bool {
	// Percorrer a string verificando se os dígitos estão em ordem crescente
	for i := 0; i < len(s)-1; i++ {
		if s[i+1] != s[i]+1 {
			return false
		}
	}
	return true
}
