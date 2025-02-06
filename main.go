package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// Executa o comando 'brew list' para pegar a lista de pacotes instalados
	cmd := exec.Command("brew", "list")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Erro ao executar o comando brew list:", err)
		return
	}

	pacotesInstalados := strings.Split(string(output), "\n")

	// Pacotes permitidos
	pacotesPermitidos := map[string]bool{
		"go":      true,
		"hashcat": true,
		"john":    true,
	}

	// Verifica pacotes não permitidos
	for _, pkg := range pacotesInstalados {
		if pkg != "" && !pacotesPermitidos[pkg] {
			fmt.Printf("Pacote não permitido encontrado: %s\n", pkg)
		}
	}

	// Verifica se todos os pacotes permitidos estão instalados
	for pkg := range pacotesPermitidos {
		encontrado := false
		for _, pacote := range pacotesInstalados {
			if pkg == pacote {
				encontrado = true
				break
			}
		}
		if !encontrado {
			fmt.Printf("Pacote permitido ausente: %s\n", pkg)
		}
	}
}
