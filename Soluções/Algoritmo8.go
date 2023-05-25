package main

import "fmt"

func main() {
	var input string
	fmt.Println("digite sua string: ")
	fmt.Scanln(&input)

	output := ""

	for i := len(input) - 1; i >= 0; i-- {
		output += string(input[i])

	}
	fmt.Println("A string invertida é: ", output)
}
