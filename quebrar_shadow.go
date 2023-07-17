package main

//A forma de uso é go run quebrar_shadow.go

import (
	"bufio"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Digite o caminho completo do arquivo: ")
	var full_path string
	fmt.Scan(&full_path)

	file, err := os.Open(full_path)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	fmt.Print("Digite o Hash completo: ")
	var comp string
	fmt.Scan(&comp)

	fmt.Print("Digite o Salt: ")
	var salt string
	fmt.Scan(&salt)

	found := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		senha := strings.TrimSpace(scanner.Text())

		hash := sha512.New()
		hash.Write([]byte(salt + senha))
		result := hex.EncodeToString(hash.Sum(nil))

		if comp == result {
			fmt.Println("Senha encontrada: ", senha)
			found = true
			break
		} else {
			fmt.Println("Testando.. ", senha)
		}
	}

	if !found {
		fmt.Println("Senha não encontrada..")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
	}
}
