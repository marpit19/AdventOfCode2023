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

	result2 := solvePuzzle2(lines)
	fmt.Println("Solution4b:", result2)
}

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

func solvePuzzle2(data []string) int {
	result := 0

	copies := make(map[int]int)

	for i := 0; i <= len(data)-1; i++ {
		ticketString := data[i]

		result += copies[i] + 1

		ticketParts := strings.Split(ticketString, ":")
		power := countWinnigNumbers(ticketParts[1])

		for c := 0; c <= power-1; c++ {
			copies[i+c+1] = copies[i+c+1] + (copies[i] + 1)
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

func countIntersection(winningNumbers []string, ticketNumbers []string) int {
	intersections := make([]string, 0)
	for _, n1 := range winningNumbers {
		for _, n2 := range ticketNumbers {
			if n1 == n2 {
				intersections = append(intersections, n2)
			}
		}
	}

	return len(intersections)
}

func countWinnigNumbers(ticketString string) int {
	ticketParts := strings.Split(ticketString, "|")
	winningNumbers := strings.Fields(ticketParts[0])
	ticketNumbers := strings.Fields(ticketParts[1])

	return countIntersection(winningNumbers, ticketNumbers)
}
