package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var CurrentVersion string = "development"

var packages = flag.String("packages", "kubectl", "Pacotes para serem instalados separados por vírgula: Kubectl")

func main() {
	flag.Parse()

	fmt.Println("Instalando: ", *packages)
	fmt.Println("Version: ", CurrentVersion)

	packageArr := strings.Split(*packages, ",")
	for _, p := range packageArr {
		switch p {
		case "kubectl":
			command, args := InstallKubectl()
			cmd := exec.Command(command, args...)

			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		default:
			fmt.Println("Pacote", p, " não suportado pelo nosso script")
		}
	}

}
