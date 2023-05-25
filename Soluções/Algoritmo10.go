package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b string
	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&a)
	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&b)

	// Converter as strings para letras minúsculas e remover os espaços em branco
	a = strings.ToLower(strings.ReplaceAll(a, " ", ""))
	b = strings.ToLower(strings.ReplaceAll(b, " ", ""))

	// Verificar se as strings têm o mesmo tamanho
	if len(a) != len(b) {
		fmt.Println("As strings não são anagramas.")
		return
	}

	// Inicializar um mapa para contar a frequência de cada letra na primeira string
	freq := make(map[rune]int)
	for _, c := range a {
		freq[c]++
	}

	// Verificar se a segunda string contém as mesmas letras da primeira string
	for _, c := range b {
		if freq[c] == 0 {
			fmt.Println("As strings não são anagramas.")
			return
		}
		freq[c]--
	}

	fmt.Println("As strings são anagramas.")
}
