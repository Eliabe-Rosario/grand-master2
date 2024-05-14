package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Digite a frase: ")
	scanner := bufio.NewScanner((bufio.NewReader(os.Stdin)))
	scanner.Scan()
	frase := scanner.Text()

	fmt.Print("Digite o número de colunas: ")
	var colunas int
	fmt.Scanln(&colunas)

	palavras := strings.Fields(frase)
	linhaAtual := ""
	for _, palavra := range palavras {
		if len(linhaAtual)+len(palavra) <= colunas {
			if linhaAtual == "" {
				linhaAtual = palavra
			} else {
				linhaAtual += " " + palavra
			}
		} else {
			fmt.Println(linhaAtual)
			linhaAtual = palavra
		}
	}
	fmt.Println(linhaAtual)
}
