package main

import "fmt"

// privada, dentro do pacote
func dummyFuntion(s string) string {
	return fmt.Sprintf("Dummy Private Function. %s", s)
}

// publica, pois começa com letra maiúscula
func DummyPublicFunction(s string) string {
	return dummyFuntion(s)
}

func main() {
	// atribuição curta de variável (declaração + atribuição)
	// short assignment (:)
	s := DummyPublicFunction("Criada pelo Ramon")
	fmt.Println(s)

	var u string
	u = dummyFuntion("Hello world again!")
	fmt.Println(u)

	// Controle de fluxo
	// if, else, switch
	var input int
	fmt.Scan(&input)
	if input == 1 {
		fmt.Println("É 1")
	} else {
		fmt.Println("Não é 1")
	}

	switch input {
	case 1:
		fmt.Println("É 1.")
	case 2:
		fmt.Println("É 2.")
	default:
		fmt.Println("Não é 1 nem 2")
	}

	// Estruturas de repetição
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
	fmt.Println("----")

	// Criando um "while true"
	running := true
	for running {
		if input == 3 {
			running = false
		}
	}
	fmt.Println("----")
}
