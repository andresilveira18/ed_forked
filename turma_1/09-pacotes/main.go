package main

import (
	f "fmt"
	"projeto/utils"
	outroutils "projeto/utils/outroUtils"
)

func main() {
	f.Println("Olá, mundo!")

	f.Println(utils.Somar(4.5, 2.2))
	f.Println(utils.Multiplicar(4.5, 2.2))
	f.Println(outroutils.Dividir(8.2, 4.1))
}
