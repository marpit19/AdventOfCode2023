package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("../../inputs/day03.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	matrix := make([][]string, 0, 10)

	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, strings.Split(line, ""))
	}

	result := solvePuzzle1(matrix)
	fmt.Println("Solution3a:", result)
}

func solvePuzzle1(data [][]string) int {
	result := 0

	for i, x := range data {
		currNo := 0
		nextToSymbol := false
		for j, y := range x {
			n, e := strconv.Atoi(y)
			if e == nil {
				currNo = currNo*10 + n
				if i == 0 && j == 0 {
					if isSymbol(data[i+1][j]) || isSymbol(data[i+1][j+1]) || isSymbol(data[i][j+1]) {
						nextToSymbol = true
					}
				} else if i == 0 && j == len(data[0])-1 {
					if isSymbol(data[i+1][j]) || isSymbol(data[i+1][j-1]) || isSymbol(data[i][j-1]) {
						nextToSymbol = true
					}
				} else if i == len(data)-1 && j == 0 {
					if isSymbol(data[i-1][j]) || isSymbol(data[i-1][j+1]) || isSymbol(data[i][j+1]) {
						nextToSymbol = true
					}
				} else if i == len(data)-1 && j == len(data[0])-1 {
					if isSymbol(data[i-1][j]) || isSymbol(data[i-1][j-1]) || isSymbol(data[i][j-1]) {
						nextToSymbol = true
					}
				} else if i == 0 {
					if isSymbol(data[i+1][j+1]) || isSymbol(data[i+1][j]) || isSymbol(data[i+1][j-1]) || isSymbol(data[i][j+1]) || isSymbol(data[i][j-1]) {
						nextToSymbol = true
					}
				} else if j == 0 {
					if isSymbol(data[i+1][j+1]) || isSymbol(data[i][j+1]) || isSymbol(data[i-1][j+1]) || isSymbol(data[i+1][j]) || isSymbol(data[i-1][j]) {
						nextToSymbol = true
					}
				} else if i == len(data)-1 {
					if isSymbol(data[i-1][j+1]) || isSymbol(data[i-1][j]) || isSymbol(data[i-1][j-1]) || isSymbol(data[i][j+1]) || isSymbol(data[i][j-1]) {
						nextToSymbol = true
					}
				} else if j == len(data[0])-1 {
					if isSymbol(data[i+1][j-1]) || isSymbol(data[i][j-1]) || isSymbol(data[i-1][j-1]) || isSymbol(data[i+1][j]) || isSymbol(data[i-1][j]) {
						nextToSymbol = true
					}
				} else {
					if isSymbol(data[i+1][j]) || isSymbol(data[i+1][j+1]) || isSymbol(data[i+1][j-1]) || isSymbol(data[i][j+1]) || isSymbol(data[i][j-1]) || isSymbol(data[i-1][j+1]) || isSymbol(data[i-1][j]) || isSymbol(data[i-1][j-1]) {
						nextToSymbol = true
					}
				}
			} else {
				if nextToSymbol {
					result += currNo
				}
				currNo = 0
				nextToSymbol = false
			}
		}
		if nextToSymbol {
			result += currNo
		}
		currNo = 0
		nextToSymbol = false
	}

	return result // Change this to the actual solution
}

func isSymbol(ch string) bool {
	if ch == "." || ch == "0" || ch == "1" || ch == "2" || ch == "3" || ch == "4" || ch == "5" || ch == "6" || ch == "7" || ch == "8" || ch == "9" {
		return false
	}
	return true
}
