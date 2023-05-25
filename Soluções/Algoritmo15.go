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

	// Substituir as vogais por asteriscos
	resultado := substituirVogais(entrada)

	// Exibir o resultado
	fmt.Println("Resultado:", resultado)
}

// Função para substituir as vogais por asteriscos
func substituirVogais(s string) string {
	// Definir as vogais
	vogais := "aeiouAEIOU"

	// Substituir as vogais por asteriscos
	for _, vogal := range vogais {
		s = strings.ReplaceAll(s, string(vogal), "*")
	}

	return s
}
