package main

import "fmt"

func findConnects(matrix [][]int, M, N, r, c int) int {
	// Se sair dos limites da matriz, retorna 0
	if r < 0 || c < 0 || r >= M || c >= N {
		return 0
	}

	// Se encontrar uma célula preenchida, explora a região
	if matrix[r][c] == 1 {
		matrix[r][c] = 0 // Marca como visitado para evitar loop infinito
		
		// O backtracking acontece aqui: ele soma a atual e retrocede conforme as ramificações terminam
		return 1 +
			findConnects(matrix, M, N, r-1, c) +
			findConnects(matrix, M, N, r+1, c) +
			findConnects(matrix, M, N, r, c-1) +
			findConnects(matrix, M, N, r, c+1) +
			findConnects(matrix, M, N, r-1, c-1) +
			findConnects(matrix, M, N, r-1, c+1) +
			findConnects(matrix, M, N, r+1, c-1) +
			findConnects(matrix, M, N, r+1, c+1)
	}

	return 0
}

func findMaxConnects(matrix [][]int, M, N int) int {
	maxConnects := 0
	for r := range M {
		for c := range N {
			maxConnects = max(maxConnects, findConnects(matrix, M, N, r, c))
		}
	}
	return maxConnects
}

func main() {
	// Exemplo exato do problema (Matriz 5x5)
	matrix := [][]int{
		{1, 1, 0, 0, 0},
		{0, 1, 1, 0, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{0, 1, 0, 1, 1},
	}

	resultado := findMaxConnects(matrix, 5, 5)

	fmt.Printf("Sample Output: %d\n", resultado)
}
