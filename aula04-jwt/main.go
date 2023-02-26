package main

import (
	"fmt"

	"github.com/StuxxNet/DescomplicandoGo/GoJwt/tokens"
)

func main() {
	fmt.Println("Gerando JWT...")
	tokens.Generate()
	fmt.Println("JWT Gerado com sucesso!")
}
