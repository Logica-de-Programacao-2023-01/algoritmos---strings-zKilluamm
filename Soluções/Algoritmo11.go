package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Println("Digite uma string: ")
	fmt.Scanln(&s)

	s = strings.ToLower(s)

	for _, v := range []string{"a", "e", "i", "o", "u"} {
		s = strings.ReplaceAll(s, v, "")

	}
	fmt.Println(s)
}
