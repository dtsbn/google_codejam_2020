package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type testCase struct {
	n          int
	activities []activity
}

type activity struct {
	num           int
	start, finish int
}

type byStartSorter []activity

func (a byStartSorter) Len() int           { return len(a) }
func (a byStartSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byStartSorter) Less(i, j int) bool { return a[i].start < a[j].start }

func (a1 *activity) Overlap(a2 *activity) bool {
	return (a2.start > a1.start && a2.start < a1.finish) || (a1.start > a2.start && a1.start < a2.finish)
}

func readInput() []testCase {
	// number of test cases
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	cases := make([]testCase, t)
	for i := 0; i < t; i++ {

		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		activities := make([]activity, n)
		for j := 0; j < n; j++ {
			scanner.Scan()
			line := strings.Split(scanner.Text(), " ")

			var activity activity
			activity.num = j
			activity.start, _ = strconv.Atoi(line[0])
			activity.finish, _ = strconv.Atoi(line[1])

			activities[j] = activity
		}

		sort.Sort(byStartSorter(activities))
		cases[i] = testCase{n, activities}
	}

	return cases
}

func makeSchedule(testCase testCase) string {
	result := make([]rune, len(testCase.activities))
	t1 := 0
	t2 := 0

	for i := 0; i < len(testCase.activities); i++ {
		if testCase.activities[i].start >= t1 {
			t1 = testCase.activities[i].finish
			result[testCase.activities[i].num] = 'C'
		} else if testCase.activities[i].start >= t2 {
			t2 = testCase.activities[i].finish
			result[testCase.activities[i].num] = 'J'
		} else {
			return "IMPOSSIBLE"
		}
	}

	return string(result)
}

func main() {
	testCases := readInput()

	results := make([]string, len(testCases))
	for i, testCase := range testCases {
		results[i] = makeSchedule(testCase)
	}

	for i, result := range results {
		fmt.Printf("Case #%d: %s\n", i+1, result)
	}
}
