//go:build homebrew

package main

import (
	"fmt"
)

func InstallKubectl() (string, []string) {
	fmt.Println("Installing kubectl")

	cmd := "brew"
	args := []string{"install", "kubectl"}
	return cmd, args
}
