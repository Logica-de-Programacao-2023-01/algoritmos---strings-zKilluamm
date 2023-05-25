package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b, c string

	fmt.Println("digite uma string: ")
	fmt.Scanln(&a)
	fmt.Println("a letra a ser substituida : ")
	fmt.Scanln(&b)
	fmt.Println("digite a nova letra: ")
	fmt.Scanln(&c)

	output := strings.ReplaceAll(a, b, c)
	fmt.Println("A nova string é: ", output)
}
