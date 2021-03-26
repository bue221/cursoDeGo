// primera clase de go
package main

import "fmt"


func main() {
	// costantes
	const pi float64 = 3.14
	const pi2 =3.14

	fmt.Println("pi ",pi)
	fmt.Println("pi2: ",pi2)

	//variables
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a,b,c,d)

	// area de un cuadrado
	const  baseCudrado = 10
	areaCuadrado := baseCudrado * baseCudrado

	fmt.Println("El area del cuadrado es: ",areaCuadrado)
}