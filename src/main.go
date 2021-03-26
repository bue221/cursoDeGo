// primera clase de go
package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) bool {
	var textReverse string
	// text a miniscula
	var textLower string = strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(textLower[i])
	}

	if textLower == textReverse {
		return true
	} else {
		return false
	}
}

func main() {
	//defer, BREAK CONTINUE
	// defer fmt.Println("Hola")
	// fmt.Println("Buena")

	//array, slices
	var array [4]int
	array[0] = 1
	array[1] = 2
	array[2] = 3

	fmt.Println(array, len(array), cap(array))

	//slices
	slice := []int{1, 2, 4, 5, 4}
	fmt.Println(slice)

	// for range

	slice2 := []string{"Andres", "aceite", "hola", "Ama"}
	for _, valor := range slice2 {
		if isPalindromo(valor) {
			fmt.Printf("La palabra %s es un palindormo", valor)
			fmt.Println("")
		} else {
			fmt.Printf("La palabra %s NO es un palindormo", valor)
			fmt.Println("")

		}
	}

}
