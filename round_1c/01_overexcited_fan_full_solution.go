package main

import (
	"fmt"
	"strconv"
)

type testCase struct {
	p    point
	path string
}

type point struct {
	x int
	y int
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

func readInput() []testCase {
	var t int
	fmt.Scanf("%d", &t)

	res := make([]testCase, t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d %d %s", &res[i].p.x, &res[i].p.y, &res[i].path)
	}
	return res
}

func calcSteps(catLocation point, path string) string {
	steps := 0
	for _, char := range path {
		if abs(catLocation.x)+abs(catLocation.y) <= steps {
			break
		}

		switch char {
		case 'N':
			catLocation.y++
		case 'S':
			catLocation.y--
		case 'E':
			catLocation.x++
		case 'W':
			catLocation.x--
		}

		steps++
	}

	if abs(catLocation.x)+abs(catLocation.y) > steps {
		return "IMPOSSIBLE"
	}
	return strconv.Itoa(steps)

}

func main() {
	tests := readInput()

	for i := 0; i < len(tests); i++ {
		fmt.Printf("Case #%d: %s\n", i+1, calcSteps(tests[i].p, tests[i].path))
	}
}
