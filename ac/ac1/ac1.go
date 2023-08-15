package main

import "fmt"

func main() {
	fmt.Println(ePrimo(7))
	fmt.Println(ePrimo(12))

	fmt.Println(fibo(5))
	fmt.Println(fibo(6))
	fmt.Println(fibo(7))

	diaSemana(4)
	diaSemana(9)

	fmt.Println(eBissexto(1995))
	fmt.Println(eBissexto(2012))
	fmt.Println(eBissexto(1900))
	fmt.Println(eBissexto(2000))
}

/*
Elabore uma função e_primo() que retorna se um número é primo ou não.
Caso o número não seja primo, liste todos os números pelos quais ele é
divisível.
*/
func ePrimo(num int) bool {
	var primo = true

	for i := 2; i < num; i++ {
		if num % i == 0 {
			primo = false
			fmt.Print(i, " ")
		}
	}

	fmt.Println("")
	return primo
}

/*
Implemente uma função fibo() que recebe um número e calcule o “n-ésimo”
elemento da série de Fibonacci. A série de Fibonacci é dada de tal forma
que um número em uma dada posição é igual à soma dos dois números anteriores.
Considere que a série é formada pela sequência 1, 1, 2, 3, …
*/
func fibo(num int) int {
	var resultado, ult, penult int
	ult, penult = 1, 1

	for i := 3; i <= num; i++ {
		resultado = ult + penult
		ult, penult = resultado, ult
	}

	return resultado
}

/*
Implementa uma função que receba um número e exiba o dia correspondente da
semana (1-Domingo, 2- Segunda, etc.), se digitar outro valor deve aparecer
valor inválido.
*/
func diaSemana(dia int) {
	switch dia {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda")
	case 3:
		fmt.Println("Terça")
	case 4:
		fmt.Println("Quarta")
	case 5:
		fmt.Println("Quinta")
	case 6:
		fmt.Println("Sexta")
	case 7:
		fmt.Println("Sábado")
	default:
		fmt.Println("Valor inválido")
	}
}

/*
Elabore uma função e_bissexto() que receba um número correspondente a um
determinado ano e em seguida informe se este ano é ou não bissexto. Um ano é
bissexto se ele é múltiplo de quatro. No entanto anos múltiplos de 100 que
não são múltiplos de 400 não são bissextos. Então 1995 não é bissexto, 2012
é bissexto, 1900 não é bissexto e 2000 é bissexto.
*/
func eBissexto(ano int) bool {
	return ano % 4 == 0 && ano % 100 != 0 || ano % 400 == 0
}
