package main

import (
	"fmt"
	"log"
	"os"
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

	return result // Change this to the actual solution
}
