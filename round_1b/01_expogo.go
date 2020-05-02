package main

import (
	"fmt"
)

type point struct {
	x float64
	y float64
}

func readInput() []point {
	var t int
	fmt.Scanf("%d", &t)

	result := make([]point, t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%f", &result[i].x)
		fmt.Scanf("%f", &result[i].y)
	}
	return result
}

func countPath(p point) string {
	right := "E"
	left := "W"
	if p.x < 0 {
		right = "W"
		left = "E"
		p.x *= -1
	}

	up := "N"
	down := "S"
	if p.y < 0 {
		up = "S"
		down = "N"
		p.y *= -1
	}

	// on the right below diagonal
	switch p {
	case point{1, 0}:
		return right
	case point{3, 0}:
		return right + right
	case point{2, 1}:
		return up + right
	case point{4, 1}:
		return down + up + right
	case point{3, 2}:
		return left + up + right
	case point{4, 3}:
		return up + up + right
	// on the right above diagonal
	case point{0, 1}:
		return up
	case point{0, 3}:
		return up + up
	case point{1, 2}:
		return right + up
	case point{1, 4}:
		return left + right + up
	case point{2, 3}:
		return down + right + up
	case point{3, 4}:
		return right + right + up
	}

	return "IMPOSSIBLE"
}

func main() {
	m := readInput()
	// fmt.Println(m)

	for i, line := range m {
		fmt.Printf("Case #%d: %s\n", i+1, countPath(line))
	}
}
