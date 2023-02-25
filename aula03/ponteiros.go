package main

import "fmt"

func main() {
	var x int64
	x = 10
	var b bool
	b = false

	// Pegando o ponteiro da variável
	// x - ponto da memória
	pointerOfX := &x
	pointerOfB := &b

	fmt.Println(x)
	fmt.Println(b)
	fmt.Println(pointerOfX)
	fmt.Println(pointerOfB)

	alfabeto := []string{"a", "b", "c", "d"}
	fmt.Println("alfabeto: ", alfabeto)

	alfabeto[0] = "z"

	primeirasLetras := alfabeto[0:2]
	fmt.Println(primeirasLetras)
	fmt.Println("alfabeto: ", alfabeto)

	alfabeto = append(alfabeto, "e")
	primeirasLetras[1] = "y"
	fmt.Println(primeirasLetras)
	fmt.Println("alfabeto: ", alfabeto)
}
