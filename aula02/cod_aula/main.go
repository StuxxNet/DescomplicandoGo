package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
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

func basicWriteIntoFile(filePath string) {
	// Criado variável pra guardar o modo do arquivo
	var perm fs.FileMode = 0744

	// Conteúdo do arquivo
	txt := []byte("Esse eh o conteudo do meu arquivo novo")

	// Escrita básica
	os.WriteFile(filePath, txt, perm)
}

func createWriteFile(filePath string, content string) {
	// Criando arquivo no SO
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
	}

	// Fechamento ao fim da função
	defer file.Close()

	// Escrevendo conteúdo no arquivo
	file.WriteString(content)
}

func createFileAppending(filePath string, content string) {
	// Abrindo arquivo, e caso ele não exista criaremos
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}

	// Fechamento ao fim da função
	defer file.Close()

	// Appendando no arquivo criado/aberto
	file.WriteString(content)
	file.WriteString("\n--------\n")
}

// Função que será chamada para cada arquivo dentro do meu
// diretório durante o walk feito pelo WalkDir
func iterateDir(path string, d fs.DirEntry, err error) error {
	fmt.Println("É diretório? ", d.IsDir())

	info, err := d.Info()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(info.Name())
	fmt.Println("\n------------")

	return nil
}

func main() {
	filepath.WalkDir(".", iterateDir)
}
