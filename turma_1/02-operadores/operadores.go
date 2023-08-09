package main

import "fmt"

func main() {
	var x, y, z = 4, 3, -2
	// aritméticos
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y) // operador módulo

	// atribuição
	z++		// z = z + 1
	z--		// z = z - 1
	z += 2	// z = z + 2
	z -= 2
	z *= 2
	z /= 2
	z %= 2

	// comparação
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x >= y)
	fmt.Println(x > y)
	fmt.Println(x <= y)
	fmt.Println(x < y)

	// lógicos
	fmt.Println(true && false) // AND
	fmt.Println(true || false) // OR
	fmt.Println(!true)		   // NOT

	// não vamos falar sobre
	// bitwise
	// memória e canal
}
