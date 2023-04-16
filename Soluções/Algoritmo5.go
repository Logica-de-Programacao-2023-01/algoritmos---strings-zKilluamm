package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	fmt.Println("Digite um número em ponto flutuante:")
	fmt.Scanln(&input)

	if _, err := strconv.ParseFloat(input, 64); err == nil {
		fmt.Println("A entrada é um número válido em ponto flutuante.")
	} else {
		fmt.Println("A entrada não é um número válido em ponto flutuante.")
	}
}
