package main

import "fmt"

func main() {
	// Array é uma coleção de dados do mesmo tipo, com tamanho definido em
	// compilação
	var filmes [5]string

	filmes[0] = "O Senhor dos Anéis"
	filmes[1] = "O Poderoso Chefão"
	fmt.Println(filmes)

	// Declaração curta
	numeros := [4]int{0, 2, 4, 6}

	fmt.Println(numeros)

	// Slices são segmentos de arrays, que podem ou não ter sido definidos
	// previamente. Eles possuem dimensão dinâmica.
	var outrosNumeros []int

	outrosNumeros = numeros[1:] // Elemento 1 (inclusive) até o final
	outrosNumeros = numeros[1:3] // Elemento 1 (inclusive) até 3 (exclusive)

	outrosNumeros[1] = 100

	fmt.Println(outrosNumeros)
	fmt.Println(numeros)

	fmt.Println("-------------")
	outrosNumeros = append(outrosNumeros, 7)

	fmt.Println("outrosNum", outrosNumeros)
	fmt.Println("num", numeros)

	fmt.Println("-------------")
	outrosNumeros = append(outrosNumeros, 9)

	outrosNumeros[1] = 200

	fmt.Println("outrosNum", outrosNumeros)
	fmt.Println("num", numeros)

	// Não precisa criar um slice a partir de um array previamente definido
	nomes := []string{"Ana", "Pedro"}

	fmt.Println(nomes)

	// Adicionar elementos no slice
	nomes = append(nomes, "João")
	fmt.Println(nomes)

	nomes = append(nomes[:1], nomes[2])
	fmt.Println(nomes)

	// Iterando sobre arrays ou slices
	fmt.Println("-------------")
	for i := 0; i < len(numeros); i++ {
		fmt.Println(numeros[i])
	}

	for indice, num := range numeros {
		fmt.Println("Índice", indice, "- valor", num)
	}

	fmt.Println("-------------")
	modificarArray(numeros) // passagem de valor dos arrays
	fmt.Println(numeros)

	outrosNumeros = numeros[1:3]
	modificarSlice(outrosNumeros) // passagem por referência dos slices
	fmt.Println(outrosNumeros)
	fmt.Println(numeros)

	slice := criarSlice()
	fmt.Println(slice)
}

func modificarArray(a [4]int) {
	a[0] = 999
	fmt.Println(a)
}

func modificarSlice(s []int) {
	s[0] = 999
}

func criarSlice() []int {
	novoSlice := []int{10, 20, 30}
	return novoSlice
}
