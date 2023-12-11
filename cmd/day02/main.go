package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("../../inputs/day02.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	result := solvePuzzle(lines)
	fmt.Println("Solution2a:", result)


}

func solvePuzzle(data []string) int {
	result := 0

	maxRed, maxGreen, maxBlue := 12, 13, 14

	for _, line := range data {
		isValid := true

		parts := strings.Split(line, ":")
		partNumber, err := strconv.Atoi(strings.TrimLeft(parts[0], "Game "))
		if err != nil {
			panic(err)
		}

		rounds := strings.Split(parts[1], ";")
		for _, round := range rounds {
			pulls := strings.Split(round, ",")
			for _, pull := range pulls {
				parts = strings.Split(strings.TrimSpace(pull), " ")
				count, err := strconv.Atoi(parts[0])
				if err != nil {
					panic(err)
				}
				color := strings.TrimSpace(parts[1])
				if color == "red" && count > maxRed {
					isValid = false
				}
				if color == "green" && count > maxGreen {
					isValid = false
				}
				if color == "blue" && count > maxBlue {
					isValid = false
				}
			}
		}

		if isValid {
			result += partNumber
		}
	}

	return result // Change this to the actual solution
}
