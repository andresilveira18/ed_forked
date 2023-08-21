package main

import (
	"fmt"
	"strings"
)

func main() {
	x := 5
	y := x

	fmt.Println(x, y) // 5 5

	x = 6
	fmt.Println(x, y) // 6 5

	z := &x // referência de x -> z recebe o endereço de x
	fmt.Println(z) // endereço de x
	fmt.Println(*z) // valor de x -> dereferência de z

	x = 7
	fmt.Println(x, *z) // 7 7
	fmt.Printf("%T\n", x) // imprime o tipo da variável

	var w *int
	fmt.Println(w)
	// fmt.Println(*w) -> retorna um panic: runtime error!

	mensagem := "Olá, mundo!" // string
	alteraMensagem(&mensagem)
	fmt.Println(mensagem) // Olá, turma!
}

func alteraMensagem(msg *string) { // msg -> ponteiro para string
	// strings.ReplaceAll(texto string, valorProcurado string, valorSubst string)
	*msg = strings.ReplaceAll(*msg, "mundo", "turma")
}
