package main

import (
	"fmt"
	"strings"
)

func main() {
	var input, OldChar, NewChar string

	fmt.Println("Digite uma string: ")
	fmt.Scanln(&input)

	fmt.Println("Digite o carácterer a ser substítuido: ")
	fmt.Scanln(&OldChar)

	fmt.Println("Digite o caractere novo carácter: ")
	fmt.Scanln(&NewChar)

	output := strings.ReplaceAll(input, OldChar, NewChar)
	fmt.Println("A nova string é: ", output)

}
