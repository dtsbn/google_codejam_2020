package main

import "fmt"

type matrix [][]int

func (m matrix) printMatrix() {
	for i := range m {
		for j := range m[i] {
			fmt.Printf("%d ", m[i][j])
		}
		fmt.Println()
	}
}

func readInput() [][]int {
	var t int
	fmt.Scanf("%d", &t)

	params := make([][]int, t)

	for i := 0; i < t; i++ {
		var n int
		var k int
		fmt.Scanf("%d", &n)
		fmt.Scanf("%d", &k)

		params[i] = []int{n, k}
	}

	return params
}

func generateMatrices(n int) int {
	// generate all matrices and count their traces
}

func makeMatrix(n, k int) (string, matrix) {
	m := make(matrix, n)

	return "POSSIBLE", m
}

func main() {
	params := readInput()

	for i, param := range params {
		result, matrix := makeMatrix(param[0], param[1])
		fmt.Printf("Case #%d: %s\n", i+1, result)
		if result == "POSSIBLE" {
			matrix.printMatrix()
		}
	}
}
