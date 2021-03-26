// primera clase de go
package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripeArgument(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a + 1
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {

	normalFunction("Buena perro")
	tripeArgument(1, 2, "Loca")
	fmt.Println(returnValue(34))
	value1, value2 := doubleReturn(2)
	fmt.Println("Values: ", value1, value2)
	// // se repite
	// fmt.Println("Hola mundo")
	// fmt.Println("Hola mundo 2")
	// fmt.Println("Hola mundo 3")

}
