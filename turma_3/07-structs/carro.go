package main

import "fmt"

type Carro struct {
	Modelo		string
	Cor			string
	Velocidade	int
	EstaLigado	bool
}

func (c *Carro) Ligar() {
	c.EstaLigado = true
}

func (c *Carro) Desligar() {
	c.EstaLigado = false
	c.Velocidade = 0
}

func (c *Carro) Acelerar(valor int) {
	c.Velocidade += valor
}

func main() {
	carro := Carro{
		Modelo:		"Celta",
		Cor:		"Preto",
		Velocidade: 0,
		EstaLigado: false,
	}

	carro.Ligar()
	fmt.Println(carro)

	carro.Acelerar(15)
	fmt.Println(carro)

	carro.Desligar()
	fmt.Println(carro)
}
