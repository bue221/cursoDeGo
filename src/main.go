// primera clase de go
package main

import "fmt"

func main() {

	// area de un cuadrado
	const baseCudrado = 10
	areaCuadrado := baseCudrado * baseCudrado

	fmt.Println("El area del cuadrado es: ", areaCuadrado)

	//variables
	x := 10
	y := 39

	//suma
	suma := x + y
	fmt.Println("suma: ", suma)
	//resta
	resta := x - y
	fmt.Println("resta: ", resta)
	//multiplicacion
	multi := x * y
	fmt.Println("Multiplicacion: ", multi)
	//division
	div := y / x
	fmt.Println("Division: ", div)
	//modulo
	mod := y % x
	fmt.Println("Modulo: ", mod)
	// Incremental el dcremental seria con --
	x++
	fmt.Println("Incremental: ", x)

	//area de un rectangulo
	const base = 40
	const altura = 30
	areaRect := base * altura
	fmt.Println("Area rectangulo: ", areaRect)

	//area de un circulo
	const radio = 4
	const pi = 3.14
	areaCir := (pi) * (radio * radio)
	fmt.Println("Area rectangulo: ", areaCir)

	//area de un trapecio
	baseUno := 6
	baseDos := 15
	alturaTra := 25

	areaTrap := ((baseUno + baseDos) * alturaTra) / 2
	fmt.Println("Area trapecio: ", areaTrap)
}
