package main

import (
	"fmt"
)

func integers() {
	fmt.Println("---------")
	fmt.Println("Integers Numbers")
	fmt.Println("---------")
	var birds int32 = 5
	var bees int16 = 10

	fmt.Println(birds)
	fmt.Println(bees)

	animals := birds + int32(bees) // Casting the tipo
	fmt.Println(animals)
}

func floats() {
	fmt.Println("---------")
	fmt.Println("Floating Numbers")
	fmt.Println("---------")
	var a float32
	var b float32

	a = 125e5 // 125 ^5
	fmt.Println(a)

	b = 12.5
	fmt.Println(b)

	c := complex(5, 2)
	fmt.Println(real(c))
	fmt.Println(imag(c))
	fmt.Println(c + complex(10, 5))
}

func strings() {
	fmt.Println("-----------")
	fmt.Println("Strings")
	fmt.Println("-----------")

	// Literal
	s := "Minha string literal"
	fmt.Println(s)
	fmt.Println(len(s))

	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]))
	}
	fmt.Println()
}

func booleans() {
	fmt.Println("-----------")
	fmt.Println("Booleans")
	fmt.Println("-----------")
	var test bool
	test = true

	fmt.Println(test)
	fmt.Println(!test)
}

func arrays() {
	fmt.Println("-----------")
	fmt.Println("Arrays")
	fmt.Println("-----------")
	caracteres := [5]string{"a", "b", "c", "d", "e"}

	fmt.Println(len(caracteres))
	fmt.Println(caracteres[len(caracteres)-1]) // Arrays começam por 0

	for i := 0; i < len(caracteres); i++ {
		fmt.Println(caracteres[i])
	}

	matriz_caracteres := [2][5]string{
		{"a", "b", "c", "d", "e"},
		{"f", "g", "h", "i", "j"},
	}

	for linha := 0; linha < len(matriz_caracteres); linha++ {
		for coluna := 0; coluna < len(matriz_caracteres[linha]); coluna++ {
			fmt.Println(matriz_caracteres[linha][coluna])
		}
	}
}

func slices() {
	fmt.Println("-----------")
	fmt.Println("Slices")
	fmt.Println("-----------")
	caracteres := []string{"a", "b", "c", "d", "e"}

	caracteres = append(caracteres, "f")

	fmt.Println(caracteres)

	// Setando tamanho inicial e final do slice pra evitar
	// alocação desnecessária no começo do programa
	tamanhoDoSlice := 10
	capacidade := 200
	caracteres = make([]string, tamanhoDoSlice, capacidade)
}

func mapas() {
	fmt.Println("-----------")
	fmt.Println("Mapas")
	fmt.Println("-----------")
	alfabeto := make(map[string]string)
	alfabeto["vogais"] = "aeiou"
	alfabeto["consoantes"] = "bcdfg"

	fmt.Println(alfabeto)
}

func estruturas() {
	fmt.Println("-----------")
	fmt.Println("Estruturas")
	fmt.Println("-----------")

	// Declarando tipo
	type Carro struct {
		Marca string
		Ano   int
	}

	// Struct usando type declarado acima
	carroTyped := Carro{
		Marca: "Gol",
		Ano:   2010,
	}

	fmt.Println(carroTyped)
	fmt.Println(carroTyped.Marca)
	fmt.Println(carroTyped.Ano)

	// Struct anônima
	carro := struct {
		Marca string
		Ano   int
	}{
		Marca: "Gol",
		Ano:   2010,
	}

	fmt.Println(carro)
	fmt.Println(carro.Marca)
	fmt.Println(carro.Ano)

}

func somaValores(num int, valores ...int) (int, int, error) {
	if num == 0 {
		return 0, 0, fmt.Errorf("Numéro não pode ser zero")
	}
	for _, valor := range valores {
		num += valor
	}

	return num, len(valores), nil
}

func main() {
	integers()
	floats()
	strings()
	booleans()
	arrays()
	slices()
	mapas()
	estruturas()

	//Nested function
	func() {
		fmt.Println("Nested function aqui!")
	}()

	// Passando parâmetro
	func(a int) {
		fmt.Println(a)
	}(5)

	// Guardando em variável
	f := func(a int) {
		fmt.Println(a)
	}
	f(10)

	// Chamando função com multiplos
	// parâmetros e multiplos retornos
	n, l, err := somaValores(5, 10, 15)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
		fmt.Println(l)
	}

}
