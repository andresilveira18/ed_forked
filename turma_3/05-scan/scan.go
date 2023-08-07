package main

import "fmt"

func main() {
	var x int
	var y float64
	var z string

	fmt.Println("Informe um valor inteiro")
	fmt.Scan(&x)
	fmt.Println(x)

	fmt.Println("Informe um float")
	fmt.Scan(&y)
	fmt.Println(y)

	fmt.Println("Informe um texto")
	fmt.Scan(&z) // pega apenas a primeira palavra, sem o espaço
	fmt.Println(z)
}
