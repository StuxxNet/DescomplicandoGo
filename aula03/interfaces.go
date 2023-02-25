package main

import "fmt"

type Veiculo interface {
	Buzinar()
}

type Carro struct {
	Marca string
	Ano   int
}

type Moto struct {
	Modelo string
	Ano    int
}

func (c *Carro) Buzinar() {
	fmt.Println("beep")
}

func (c *Moto) Buzinar() {
	fmt.Println("bop bop bop")
}

func buzina(v Veiculo) {
	v.Buzinar()
}

func main() {
	carro := Carro{
		Marca: "Chevrolet",
		Ano:   2010,
	}
	moto := Moto{
		Modelo: "Cinquentinha",
		Ano:    1992,
	}
	frota := (make([]Veiculo, 2))
	frota[0] = &carro
	frota[1] = &moto

	// Carro e Moto são veículos
	// porque ambos implementam o método Buzinar
	// que está definido na interface. Quando temos uma
	// funcão que receve um valor do tipo da interface
	// a passagem de parâmetro exige um ponteiro
	for i := 0; i < len(frota); i++ {
		buzina(frota[i])
	}
}
