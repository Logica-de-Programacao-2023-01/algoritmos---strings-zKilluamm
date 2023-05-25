package main

import (
	"fmt"
	"strings"
)

func main() {
	// Solicitar uma string ao usuário
	fmt.Print("Digite uma string: ")
	var entrada string
	fmt.Scanln(&entrada)

	// Obter as letras únicas da string
	letrasUnicas := obterLetrasUnicas(entrada)

	// Imprimir as letras únicas
	fmt.Println("Letras únicas:", letrasUnicas)
}

// Função para obter as letras únicas de uma string
func obterLetrasUnicas(s string) string {
	letras := ""
	for _, char := range s {
		letra := string(char)
		if strings.Count(s, letra) == 1 {
			letras += letra
		}
	}
	return letras
}
