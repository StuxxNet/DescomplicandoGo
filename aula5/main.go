package main

import (
	"log"
	"os/exec"
)

func main() {
	command, args := InstallKubectl()
	cmd := exec.Command(command, args...)

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
