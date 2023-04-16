package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	output := strings.ReplaceAll(input, " ", "")
	fmt.Println("string sem espaços", output)
}
