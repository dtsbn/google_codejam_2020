package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() []string {
	// get number of input lines
	var n int
	fmt.Scanf("%d", &n)

	lines := make([]string, n)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		lines[i] = scanner.Text()
	}

	return lines
}

// convert string of ints to []int
func makeNums(line string) []int {
	strs := strings.Split(line, "")
	nums := make([]int, len(strs))

	for i, str := range strs {
		nums[i], _ = strconv.Atoi(str)
	}

	return nums
}

func makeLineWithBrackets(nums []int) string {
	var bracketsLine string

	counter := 0
	for i := 0; i < len(nums); i++ {
		// fmt.Printf("i: %d, counter: %d, nums[i]: %d\n", i, counter, nums[i])
		for counter < nums[i] {
			bracketsLine += "("
			counter++
		}

		for counter > nums[i] {
			bracketsLine += ")"
			counter--
		}

		bracketsLine += strconv.Itoa(nums[i])
	}

	// add closing brackets
	for counter > 0 {
		bracketsLine += ")"
		counter--
	}

	return bracketsLine
}

func main() {
	lines := readInput()
	numsLines := make([][]int, len(lines))

	for i, line := range lines {
		numsLines[i] = makeNums(line)
	}

	bracketsLines := make([]string, len(lines))
	for i, line := range numsLines {
		bracketsLines[i] = makeLineWithBrackets(line)
	}

	// fmt.Println("Check if lines are correct:")
	// for _, line := range lines {
	// 	fmt.Println(line)
	// }

	// fmt.Println("Check if strings converted to ints are correct:")
	// for _, line := range numsLines {
	// 	fmt.Println(line)
	// }

	// fmt.Println("Check if lines with brackets are correct:")
	// for _, line := range bracketsLines {
	// 	fmt.Println(line)
	// }

	for i, line := range bracketsLines {
		fmt.Printf("Case #%d: %s\n", i+1, line)
	}

}

/* Test example from Google:
4
0000	0000
101		(1)0(1)
111000	(111)000
1		(1)
*/
