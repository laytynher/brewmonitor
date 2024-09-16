package main

import (
    "fmt"
    "os/exec"
    "strings"
)

func main() {
    // Executa o comando 'brew list' para obter a lista de pacotes instalados
    cmd := exec.Command("brew", "list")
    output, err := cmd.Output()
    if err != nil {
        fmt.Println("Erro ao executar brew list:", err)
        return
    }

    installedPackages := strings.Split(string(output), "\n")

    // Pacotes permitidos
    allowedPackages := map[string]bool{
        "go":      true,
        "hashcat": true,
        "john":    true,
    }

    // Verifica se há pacotes não permitidos
    for _, pkg := range installedPackages {
        if pkg != "" && !allowedPackages[pkg] {
            fmt.Printf("Pacote não permitido detectado: %s\n", pkg)
        }
    }

    // Verifica se todos os pacotes permitidos estão instalados
    for pkg := range allowedPackages {
        found := false
        for _, installedPkg := range installedPackages {
            if pkg == installedPkg {
                found = true
                break
            }
        }
        if !found {
            fmt.Printf("Pacote permitido não encontrado: %s\n", pkg)
        }
    }
}
