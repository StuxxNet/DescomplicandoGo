//go:build homebrew

package main

import (
	"fmt"
)

func InstallKubectl() (string, []string) {
	fmt.Println("Installing kubectl")

	cmd := "curl"
	args := []string{"-LO", "https://dl.k8s.io/release/v1.26.2/bin/linux/amd64/kubectl"}
	return cmd, args
}
