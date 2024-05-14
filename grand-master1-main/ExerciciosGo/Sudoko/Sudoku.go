package main

import (
	"fmt"
)

const tamanhoGrid = 9

func éSudokuVálido(tabuleiro [][]int) bool {
	for i := 0; i < tamanhoGrid; i++ {
		mapaLinha := make(map[int]bool)
		mapaColuna := make(map[int]bool)
		for j := 0; j < tamanhoGrid; j++ {
			if tabuleiro[i][j] != 0 {
				if mapaLinha[tabuleiro[i][j]] {
					return false
				}
				mapaLinha[tabuleiro[i][j]] = true
			}
			if tabuleiro[j][i] != 0 {
				if mapaColuna[tabuleiro[j][i]] {
					return false
				}
				mapaColuna[tabuleiro[j][i]] = true
			}
		}
	}

	for i := 0; i < tamanhoGrid; i += 3 {
		for j := 0; j < tamanhoGrid; j += 3 {
			if !éRegiãoVálida(tabuleiro, i, j) {
				return false
			}
		}
	}

	return true
}

func éRegiãoVálida(tabuleiro [][]int, linha, coluna int) bool {
	mapaRegião := make(map[int]bool)
	for i := linha; i < linha+3; i++ {
		for j := coluna; j < coluna+3; j++ {
			if tabuleiro[i][j] != 0 {
				if mapaRegião[tabuleiro[i][j]] {
					return false
				}
				mapaRegião[tabuleiro[i][j]] = true
			}
		}
	}
	return true
}

func main() {
	tabuleiro := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	if éSudokuVálido(tabuleiro) {
		fmt.Println("O tabuleiro de Sudoku é válido.")
	} else {
		fmt.Println("O tabuleiro de Sudoku contém valores incorretos.")
	}
}
