// primera clase de go
package main

import "fmt"

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("No es cierto")
	}

	if valor1 == 1 || valor2 == 2 {
		fmt.Println("Cierto")
	} else {
		fmt.Println("Falso")
	}
}
