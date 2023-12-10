package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("../../inputs/day01.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	result := solvePuzzle(lines)
	fmt.Println("Solution:", result)
}

func solvePuzzle(data []string) int {
	result := 0

	for _, line := range data {
		result += (10 * getFirstNumber(line)) + getLastNumber(line)
	}

	return result // Change this to the actual solution
}

func getFirstNumber(line string) int {
	for _, c := range line {
		if digit, err := strconv.Atoi(string(c)); err == nil {
			return digit
		}
	}
	return -1
}

func getLastNumber(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		c := line[i]

		if digit, err := strconv.Atoi(string(c)); err == nil {
			return digit
		}
	}
	return -1
}
