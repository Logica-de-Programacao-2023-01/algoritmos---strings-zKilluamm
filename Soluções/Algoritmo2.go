package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Digite uma string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	result := strings.ReplaceAll(input, " ", "")
	fmt.Printf("A string sem espaços em branco é: %s\n", result)
}
