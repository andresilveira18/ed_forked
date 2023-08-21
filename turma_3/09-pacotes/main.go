package main

import (
	"fmt"
	"projetinho/utils"
	"projetinho/utils/outroutils"
)

func main() {
	fmt.Println("Olá, mundo!")

	fmt.Println(utils.Somar(4.3, 6.2))
	fmt.Println(utils.Subtrair(4.3, 6.2))
	fmt.Println(utils.Multiplicar(4.3, 6.2))
	fmt.Println(outroutils.Dividir(5.4, 2.0))
}
