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
	fmt.Println("Solution 1a:", result)

	result2 := solvePuzzle2(lines)
	fmt.Println("Solution 1b:", result2)
}

func solvePuzzle(data []string) int {
	result := 0

	for _, line := range data {
		result += (10 * getFirstNumber(line)) + getLastNumber(line)
	}

	return result // Change this to the actual solution
}

func solvePuzzle2(data []string) int {
	result := 0

	for _, line := range data {
		line = replaceStringByInt(line)
		result += (10 * getFirstNumber(line)) + getLastNumber(line)
	}

	return result
}

func replaceStringByInt(line string) string {
	stringInt := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	for idx, s := range stringInt {
		conv := strconv.Itoa(idx + 1)
		firstChar := string(s[0])
		lastChar := string(s[len(s)-1])
		replace := firstChar + conv + lastChar
		line = strings.ReplaceAll(line, s, replace)
	}
	// strings.ReplaceAll to find all occurrences of the current string s in the line and replaces them with
	// the replace string. This updates the line variable with the replacements

	return line
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
