// primera clase de go
package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepa"] = 14

	fmt.Println(m)

	//recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//encontrar un valor
	value := m["jose"]
	fmt.Println(value)
}
