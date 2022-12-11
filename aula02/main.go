package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func checkSize(filePath string) {
	// multiplos retornos
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Fecha o arquivo ao fim da execução da função
	defer file.Close()

	// Printando arquivo
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(stat.Size())

	// Fechando o arquivo
	file.Close()
}

func basicFileRead(filePath string) {
	// Todo conteúdo em uma única variável
	// Ótimo para leitura de arquivos pequenos, pois tudo na mesma
	// variável na memória
	file, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("File content: ", string(file))
}

func readViaIoScanner(filePath string) {
	// Leitura via buffer. A gente joga o conteúdo por partes
	// e vai lendo quando faz o scan
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Fecha o arquivo ao fim da execução da função
	defer file.Close()

	// Todo file é um Reader pois implementa a interface Reader
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Scanner pode terminar em erro, então bom fazer check pois fim
	err = scanner.Err()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func readByteByByte(filePath string) {
	// Leitura byte por byte do arquivo
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Fecha o arquivo ao fim da execução da função
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		b, err := reader.ReadByte()
		// Printa erro só se for diferente do fim do arquivo
		if err != nil {
			if err != io.EOF {
				fmt.Println(err.Error())
			}
			break
		}

		// Printando sem ln pra não quebrar a cada letra
		// Dependendo do encoding cada letra por ser um byte ou mais
		fmt.Print(string(b))
	}
}

func main() {
	readViaIoScanner("TA_PRECO_MEDICAMENTO.CSV")
}
