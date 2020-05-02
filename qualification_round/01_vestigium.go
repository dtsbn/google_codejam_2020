package main

import (
	"fmt"
)

type matrix [][]int

func (m matrix) printMatrix() {
	for i := range m {
		for j := range m[i] {
			fmt.Printf("%d ", m[i][j])
		}
		fmt.Println()
	}
}

func readInput() []matrix {
	// num of test cases
	var t int
	fmt.Scanf("%d", &t)

	matrices := make([]matrix, t)
	for i := 0; i < t; i++ {
		// matrix size
		var n int
		fmt.Scanf("%d", &n)

		matrices[i] = make(matrix, n)
		for j := range matrices[i] {
			matrices[i][j] = make([]int, n)
		}

		for j := range matrices[i] {
			for k := range matrices[i][j] {
				fmt.Scanf("%d", &matrices[i][j][k])
			}
		}
	}

	return matrices
}

func countMatrix(m matrix) (int, int, int) {
	var trace int
	var repeatedRows int
	var repeatedColumns int

	n := len(m)

	for i := range m {
		trace += m[i][i]
	}

	for i := 0; i < n; i++ {
		rowMap := make(map[int]int)
		for j := 0; j < n; j++ {
			rowMap[m[i][j]]++
		}
		if len(rowMap) != n {
			repeatedRows++
		}
	}

	for j := 0; j < n; j++ {
		columnMap := make(map[int]int)
		for i := 0; i < n; i++ {
			columnMap[m[i][j]]++
		}
		if len(columnMap) != n {
			repeatedColumns++
		}
	}

	return trace, repeatedRows, repeatedColumns
}

func main() {
	matrices := readInput()

	for i, matrix := range matrices {
		trace, rows, columns := countMatrix(matrix)
		fmt.Printf("Case #%d: %d %d %d\n", i+1, trace, rows, columns)
	}
}
