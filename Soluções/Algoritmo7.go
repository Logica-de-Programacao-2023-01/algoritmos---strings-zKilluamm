package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite a string: ")
	scanner.Scan()
	PossuiNum := false
	palavras := scanner.Text()

	for _, i := range palavras {
		if unicode.IsDigit(i) {
			PossuiNum = true
			break
		}

	}
	if PossuiNum {
		fmt.Println("Possui número!")

	} else {
		fmt.Println("Não possui número!")
	}
}
