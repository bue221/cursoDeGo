// primera clase de go
package main

import "fmt"

func main() {
	modulo := 5 % 2
	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}
	//sin condicion
	valor := 200
	switch {
	case valor > 100:
		fmt.Println("El valor es mayor que 100")
	case valor < 100:
		fmt.Println("El valor es menor que 100")
	}
}
