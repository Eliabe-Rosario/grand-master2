package main

import (
	"fmt"
	"strings"
)

var codonMap = map[string]string{
	"AUG": "Metionina",
	"UUU": "Fenilalanina", "UUC": "Fenilalanina",
	"UUA": "Leucina", "UUG": "Leucina",
	"UCU": "Serina", "UCC": "Serina", "UCA": "Serina", "UCG": "Serina",
	"UAU": "Tirosina", "UAC": "Tirosina",
	"UGU": "Cisteína", "UGC": "Cisteína",
	"UGG": "Triptofano",
	"UAA": "PARADA", "UAG": "PARADA", "UGA": "PARADA",
}

func traduzirRNA(rna string) []string {
	var proteinas []string

	for i := 0; i < len(rna)-2; i += 3 {
		codon := rna[i : i+3]
		proteina, encontrado := codonMap[codon]
		if !encontrado || proteina == "PARADA" {
			break
		}
		proteinas = append(proteinas, proteina)
	}

	return proteinas
}

func main() {
	rna := "AUGUUUUCUUAAAUG"
	proteinas := traduzirRNA(rna)
	fmt.Println("Proteínas:", strings.Join(proteinas, ", "))
}
