package main

import (
	"fmt"

	"strings"
)

func main() {

	var input string
	fmt.Println("Digite a string: ")
	fmt.Scanln(&input)

	output := strings.ToUpper(input)
	fmt.Println("String convertida: ", output)

}
