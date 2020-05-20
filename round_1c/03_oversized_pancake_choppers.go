package main

import (
	"fmt"
	"sort"
)

type testCase struct {
	n      int         // # of slices
	d      int         // # of diners
	angles map[int]int // internal angles for each slice
}

func readInput() []testCase {
	var t int
	fmt.Scanf("%d", &t)

	res := make([]testCase, t)

	for i := 0; i < t; i++ {
		var n, d int
		fmt.Scanf("%d %d", &n, &d)

		angles := make(map[int]int)
		for j := 0; j < n; j++ {
			var a int
			fmt.Scanf("%d", &a)
			angles[a]++
		}

		res[i].n = n
		res[i].d = d
		res[i].angles = angles
	}
	return res
}

/*
func cutToThreeEqual(angles map[int]int) bool {
	// если есть 2 одинаковых куска и можно найти третий, который больше
	// чем первые два => от него можно отрезать нужный кусок
	for k, v := range angles {
		if v == 2 {
			for m, n := range angles {
				if m != k && n >= v {
					return true
				}
			}
		}
	}

	// если есть кусок, в два раза больше текущего
	for k := range angles {
		if _, ok := angles[k*2]; ok {
			return true
		}
	}

	return false
}

func calcCuts(test testCase) int {
	maxEquals := 0
	for _, v := range test.angles {
		if v > maxEquals {
			maxEquals = v
		}
	}

	if test.d == 2 {
		// нужно либо найти 2 одинаковых куска
		if maxEquals > 1 {
			return 0
		}
		// либо разрезать любой на 2
		return 1
	} else if test.d == 3 {
		// нужно либо найти 3 одинаковых куска
		if maxEquals >= 3 {
			return 0
		}
		if cutToThreeEqual(test.angles) == true {
			return 1
		}
	}
	return 2
}
*/

func calculateCuts(test testCase) int {
	for _, v := range test.angles {
		if v >= test.d {
			return 0
		}
	}

	if test.d == 2 {
		return 1
	}

	res := make([]int, len(test.angles))

	// creating sorted list of slices size
	anglesKeys := make([]int, len(test.angles))
	i := 0
	for k := range test.angles {
		anglesKeys[i] = k
		i++
	}
	sort.Ints(anglesKeys)

	// fmt.Println("test.angles: ", test.angles)
	// fmt.Println("anglesKeys: ", anglesKeys)

	for i := 0; i < len(anglesKeys); i++ {
		// r := test.d // init
		r := 0

		// fmt.Println("Init r: ", r)
		// fmt.Println("Res 1: ", res)

		pieces := 0
		for j := i; j < len(anglesKeys); j++ {
			// fmt.Println("j: ", j)
			// fmt.Println("anglesKeys[i]: ", anglesKeys[i])
			// fmt.Println("anglesKeys[j]: ", anglesKeys[j])

			if anglesKeys[j]%anglesKeys[i] == 0 {
				// r -= test.angles[anglesKeys[j]]
				r += (anglesKeys[j]/anglesKeys[i] - 1) * test.angles[anglesKeys[j]]
			}

			pieces += anglesKeys[j] / anglesKeys[i]
			// fmt.Println("Pieces: ", pieces)
			// fmt.Println("r: ", r)
			if pieces > test.d {
				r -= (pieces - test.d - 1)
				break
			}
		}

		if pieces < test.d {
			for j := i; j < len(anglesKeys); j++ {
				res[j] = test.d - 1
			}
			break
		} else {
			// fmt.Println("Final r: ", r)
			res[i] = r
		}

		// fmt.Println("Res 4: ", res)
		// fmt.Println()
	}
	sort.Ints(res)

	// fmt.Println("Res 5: ", res)

	return res[0]
}

func main() {
	tests := readInput()
	//fmt.Println("Tests: ", tests)

	for i := 0; i < len(tests); i++ {
		// fmt.Printf("Case #%d: %d\n", i+1, calcCuts(tests[i]))
		fmt.Printf("Case #%d: %d\n", i+1, calculateCuts(tests[i]))
	}
}

// TODO: написать тест-кейсов
