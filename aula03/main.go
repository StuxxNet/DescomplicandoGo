package main

import (
	"encoding/json"
	"fmt"
)

type Veiculo interface {
	Buzinar()
}

type Moto struct {
	Modelo string
	Ano    int
}

func (v *Moto) Buzinar() {
	fmt.Println("Boooop!")
}

type Carro struct {
	Marca string
	Ano   int
}

func (c *Carro) Buzinar() {
	fmt.Println("Beeep!")
}

func main() {
	carro := Carro{
		Marca: "Chevrolet",
		Ano:   2020,
	}
	moto := Moto{
		Modelo: "CB1000",
		Ano:    2022,
	}
	frota := []Veiculo{&carro, &moto}

	data, err := json.Marshal(frota)
	if err != nil {
		fmt.Print(err.Error())
		return
	} else {
		fmt.Println(string(data))
	}

	s := "[{\"Marca\":\"Chevrolet\",\"Ano\":2020}]"
	carros := []Carro{}
	err = json.Unmarshal([]byte(s), &carros)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(carros)
	}
}
