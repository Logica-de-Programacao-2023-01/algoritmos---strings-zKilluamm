package main

import (
	"fmt"
	"strings"
)

func main() {
	var string1, string2 string
	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&string1)

	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&string2)

	if strings.Compare(string1, string2) == 0 {
		fmt.Println("As strings são iguais. ")
	} else {
		fmt.Println("As strings são diferentes. ")
	}

}
