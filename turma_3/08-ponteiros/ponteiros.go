package main

import (
	"fmt"
	"strings"
)

type Pessoa struct {
	Nome  string
	Email string
}

func (p *Pessoa) alteraEmail(novoEmail string) {
	p.Email = novoEmail
}

func adicionaPessoa(p Pessoa, a *[5]Pessoa) {
	for ind, pessoa := range a {
		if (pessoa == Pessoa{}) {
			a[ind] = p
			break
		}
	}
}

func main() {
	var pessoas [5]Pessoa
	var p1 = Pessoa{
		Nome: "aaa",
		Email: "bbb",
	}

	fmt.Println(pessoas)
	adicionaPessoa(p1, &pessoas)
	fmt.Println(pessoas)

	x := 5
	y := x

	fmt.Println(x, y) // 5 5

	x = 6
	fmt.Println(x, y) // 6 5

	z := &x         // z recebe a referência de x
	fmt.Println(z)  // endereço na memória em que x está armazenado
	fmt.Println(*z) // dereferência -> retorna o valor que está sendo apontado por z

	x = 7
	fmt.Println(x, *z) // 7 7

	var w *int
	fmt.Println(w) // nil
	// fmt.Println(*w) -> dá um erro!

	mensagem := "Olá, mundo!"
	alteraMensagem(&mensagem)
	fmt.Println(mensagem) // Olá, turma!

	a := &z

	fmt.Println(z)   // end x
	fmt.Println(a)   // end z
	fmt.Println(*a)  // end x
	fmt.Println(**a) // valor x
}

func alteraMensagem(msg *string) {
	// strings.ReplaceAll(texto string, valorProcurado string, valorSubstituir string)
	// Montar uma função que substitui "mundo" por "turma"
	*msg = strings.ReplaceAll(*msg, "mundo", "turma")
}