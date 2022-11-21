package main

import "fmt"

func pegarInput() int {
	var numero int
	fmt.Scanf("%d", &numero)
	return numero
}

func somar(i int, j int) int {
	return i + j
}

func subtrair(i int, j int) int {
	return i - j
}

func dividir(i int, j int) int {
	if j == 0 {
		return 0
	}
	return i / j
}

func multiplicar(i int, j int) int {
	return i * j
}

func main() {

	primeiroNumero := pegarInput()
	segundoNumero := pegarInput()

	fmt.Printf("Operações com %d e %d\n", primeiroNumero, segundoNumero)
	fmt.Printf("O valor da soma é: %d\n", somar(primeiroNumero, segundoNumero))
	fmt.Printf("O valor da subtração é: %d\n", subtrair(primeiroNumero, segundoNumero))
	fmt.Printf("O valor da multiplicação é: %d\n", multiplicar(primeiroNumero, segundoNumero))
	fmt.Printf("O valor da divisão é: %d\n", dividir(primeiroNumero, segundoNumero))

}
