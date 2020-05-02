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

func readInput() []testCase {
	var t int
	fmt.Scanf("%d", &t)

	res := make([]testCase, t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d %d %s", &res[i].p.x, &res[i].p.y, &res[i].path)
	}
	return res
}

func catStep(catLocation *point, char rune) {
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
}

func calcSteps(test testCase) string {
	myLocation := point{0, 0}
	catLocation := test.p
	steps := 0

	for _, char := range test.path {
		if myLocation == catLocation {
			break
		}
		steps++
		catStep(&catLocation, char)
		if myLocation == catLocation {
			break
		}

		if myLocation.x != catLocation.x {
			if myLocation.x < catLocation.x {
				myLocation.x++
			} else {
				myLocation.x--
			}
		} else {
			if myLocation.y < catLocation.y {
				myLocation.y++
			} else {
				myLocation.y--
			}
		}
	}

	if myLocation != catLocation {
		return "IMPOSSIBLE"
	}
	return strconv.Itoa(steps)
}

func main() {
	tests := readInput()

	for i := 0; i < len(tests); i++ {
		fmt.Printf("Case #%d: %s\n", i+1, calcSteps(tests[i]))
	}
}
