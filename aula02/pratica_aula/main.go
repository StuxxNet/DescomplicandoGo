package main

import (
	"fmt"
	"os"
)

func writeLineToFile(filePath string, content string) {
	// Abrindo arquivo - Falha caso ele não exista.
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Fechamento ao fim da execução
	defer file.Close()

	// Gravação guardando erro com _ pra
	// ignorar primeiro erro
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	// Novo remédio
	var newMedicine string = "NEO-SINEFRINA;13.078.518/0001-90;FUNDA��O BAIANA DE PESQ. CIENTIFICA E DESENV.  TECNOLOGICO, FORNECIMENTO E DISTRIBUI��O DE MEDICAMENTOS-BAHIAFARMA;5,44018E+14;1988300050018;7,8986E+12;    -     ;    -     ;BAHIAFARMA INSULINA HUMANA R;100 UI/ML SOL INJ CT FA VD INC X 10 ML;A10C1 - INSULINA HUMANA E AN�LOGOS, DE A��O R�PIDA;    -     ;Regulado;37,95;37,95;;;;;;;;;52,46;;;;;;;;;N�o;Sim;Sim;Sim;;Positiva;Sim;Tarja Vermelha (**)"

	// Gravando no CSV
	writeLineToFile("TA_PRECO_MEDICAMENTO.CSV", newMedicine)
}
