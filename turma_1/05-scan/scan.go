package main

import "fmt"

func main() {
	var x int
	var y float64
	var msg string

	fmt.Println("Informe um número:")
	fmt.Scanln(&x)

	fmt.Println("Informe um float:")
	fmt.Scan(&y)

	fmt.Println("Informe um texto:")
	fmt.Scan(&msg)

	fmt.Println(x, y, msg)
}
