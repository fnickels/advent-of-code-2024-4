package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(fileName string) (State, error) {

	state := State{
		lines: [][]rune{},
	}

	file, err := os.Open(fileName)
	if err != nil {
		return state, fmt.Errorf("error opening file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineslice := []rune{}

		for _, char := range line {
			lineslice = append(lineslice, char)
		}

		state.lines = append(state.lines, lineslice)
	}

	return state, nil
}
