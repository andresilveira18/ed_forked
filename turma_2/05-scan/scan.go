package main

import "fmt"

func main() {
	var x int
	var y float64

	fmt.Println("Informe um número: ")
	fmt.Scan(&x)
	fmt.Println(x)

	fmt.Println("Informe um float: ")
	fmt.Scan(&y)
	fmt.Println(y)
}