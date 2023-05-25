package main

import (
	"fmt"
)

func main() {
	// Solicitar uma string ao usuário
	fmt.Print("Digite uma sequência numérica: ")
	var entrada string
	fmt.Scanln(&entrada)

	// Verificar se a string é uma sequência numérica decrescente
	if ehSequenciaDecrescente(entrada) {
		fmt.Println("A sequência é numérica decrescente!")
	} else {
		fmt.Println("A sequência não é numérica decrescente.")
	}
}

// Função para verificar se uma string é uma sequência numérica decrescente
func ehSequenciaDecrescente(s string) bool {
	// Percorrer a string verificando se os dígitos estão em ordem decrescente
	for i := 0; i < len(s)-1; i++ {
		if s[i+1] != s[i]-1 {
			return false
		}
	}
	return true
}
