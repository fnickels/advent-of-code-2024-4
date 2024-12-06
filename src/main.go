package main

import (
	"fmt"
	"os"
)

type Result struct {
	part1 int
	part2 int
}

type State struct {
	lines [][]rune
}

// main is the entry point of the Game of Life application. It reads and validates the input,
// processes the locations according to the game rules, and displays the results.
func main() {

	// read & validate input
	state, err := readInput("input-file.txt")
	if err != nil {
		fmt.Printf("error reading input: %v\n", err)
		os.Exit(1)
	}

	// process state
	result := process(state)

	// display results
	display(result)
}

func process(state State) Result {
	// create a new state to hold the results
	result := Result{}

	//for _, line := range state.lines {
	//	fmt.Println(line)
	//}

	match := []rune{'X', 'M', 'A', 'S'}

	result.part1 = countMatches(state, match)

	return result
}

func display(result Result) {
	fmt.Printf("The part 1 result is: %d\n", result.part1)
	fmt.Printf("The part 2 result is: %d\n", result.part2)
}

func countMatches(state State, match []rune) int {
	count := 0
	for i, line := range state.lines {
		for j := range line {
			count += matchCount(state, i, j, match)
		}
	}
	return count
}

func matchCount(state State, i, j int, match []rune) int {

	count := 0

	if i == 5 && j == 114 {
		count = 0 // set break point here
	}

	mLen := len(match)
	reverseMatch := make([]rune, mLen)
	for z := mLen - 1; z >= 0; z-- {
		reverseMatch = append(reverseMatch, match[z])
	}

	// horizontal match
	if j+mLen <= len(state.lines[i]) {
		testSlice := state.lines[i][j : j+mLen]
		if matchRuneSlice(testSlice, match) {
			count++
		}
		if matchRuneSlice(testSlice, reverseMatch) {
			count++
		}
	}

	// vertical match
	if i+mLen <= len(state.lines) {
		testSlice := make([]rune, mLen)
		for z := i; z < i+mLen; z++ {
			testSlice = append(testSlice, state.lines[z][j])
		}
		if matchRuneSlice(testSlice, match) {
			count++
		}
		if matchRuneSlice(testSlice, reverseMatch) {
			count++
		}
	}

	// diagonal match
	if i+mLen <= len(state.lines) && j+mLen <= len(state.lines[i]) {
		testSlice := make([]rune, mLen)
		for z := i; z < i+mLen; z++ {
			for y := j; y < j+mLen; y++ {
				testSlice = append(testSlice, state.lines[z][j])
			}
		}
		if matchRuneSlice(testSlice, match) {
			count++
		}
		if matchRuneSlice(testSlice, reverseMatch) {
			count++
		}
	}

	// reverse diagonal match
	if i+mLen <= len(state.lines) && j-mLen > 0 {
		testSlice := make([]rune, mLen)
		for z := i; z < i+mLen; z++ {
			for y := j; y > j-mLen; y-- {
				testSlice = append(testSlice, state.lines[z][j])
			}
		}
		if matchRuneSlice(testSlice, match) {
			count++
		}
		if matchRuneSlice(testSlice, reverseMatch) {
			count++
		}
	}

	if count > 0 {
		fmt.Printf("Match count for %d, %d  ==  %d\n", i, j, count)
		displayTargets(state, i, j, len(match))
	}

	return count
}

func matchRuneSlice(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func displayTargets(state State, i, j, mLen int) {
	// fmt.Printf("Displaying targets for %d, %d\n", i, j)

	for z := i; z < i+mLen; z++ {
		for y := j - mLen; y < j+mLen; y++ {
			display := " "
			if y < 0 || y >= len(state.lines[0]) || z >= len(state.lines) {
				display = "_"
			} else if (z == i && y >= j) || // horizontal
				(y == j) { // vertical
				display = string(state.lines[z][y])
			} else {
				zDiff := z - i
				yDiff := y - j
				if zDiff == yDiff || zDiff == -yDiff { // diagonal
					display = string(state.lines[z][y])
				}
			}
			fmt.Printf("%v", display)
		}
		fmt.Printf("\n")
	}
}
