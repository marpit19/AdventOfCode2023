package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("../../inputs/day04.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	result := solvePuzzle(lines)
	fmt.Println("Solution:", result)
}

// Card   1: 99 71 95 70 36 79 78 84 31 10 |  5 45 54 83  3 38 89 35 80 49 76 15 63 20 21 94 65 55 44  4 75 56 85 92 90

func solvePuzzle(data []string) int {
	result := 0

	for _, card := range data {

		winningNumbers := make(map[int]struct{})
		givenNumbers := make(map[int]struct{})
		parts := strings.Split(card, "|")

		winningString := strings.TrimSpace(parts[0])
		givenString := strings.TrimSpace(parts[1])

		winNums := extractNumbers(winningString)
		for _, n := range winNums {
			winningNumbers[n] = struct{}{}
		}

		givenNums := extractNumbers(givenString)
		for _, num := range givenNums {
			givenNumbers[num] = struct{}{}
		}

		commonNos := 0
		for num := range winningNumbers {
			if _, exists := givenNumbers[num]; exists {
				commonNos += 1
			}
		}

		if commonNos > 0 {
			x := math.Pow(2, float64(commonNos-1))
			result += int(x)
		}

	}

	return result // Change this to the actual solution
}

func extractNumbers(s string) []int {
	numStrs := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' '
	})

	nums := make([]int, 0)
	for _, numStr := range numStrs {
		num, err := strconv.Atoi(numStr)
		if err == nil {
			nums = append(nums, num)
		}
	}
	return nums
}
