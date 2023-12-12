package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	result2 := solvePuzzle2(matrix)
	fmt.Println("Solution3b:", result2)
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

func solvePuzzle2(data [][]string) int {
	res := 0
	for i, u := range data {
		for j, v := range u {
			if v == "*" {
				a1 := int(math.Max(float64(i-1), 0))
				a2 := int(math.Min(float64(i+1), float64(len(data)-1)))
				b1 := int(math.Max(float64(j-3), 0))
				b2 := int(math.Min(float64(j+3), float64(len(data[0])-1)))
				subMatrix := make([][]string, 0, 3)
				for n := 0; n <= a2-a1; n++ {
					tmp := make([]string, 0, 5)
					for m := 0; m <= b2-b1; m++ {
						if isSymbol(data[a1+n][b1+m]) {
							tmp = append(tmp, ".")
						} else {
							tmp = append(tmp, data[a1+n][b1+m])
						}
					}
					subMatrix = append(subMatrix, tmp)
				}
				subMatrix[1][3] = "*"
				res += computeSubMatrix(subMatrix)
			}
		}
	}

	return res
}

func computeSubMatrix(matrix [][]string) int {
	res := 1
	conteur := 0
	for i, u := range matrix {
		currentNumber := 0
		nextToSymbol := false
		for j, v := range u {
			number, e := strconv.Atoi(v)
			if e == nil {
				currentNumber = currentNumber*10 + number
				if i == 0 && j == 0 {
					if isSymbol(matrix[i+1][j]) || isSymbol(matrix[i+1][j+1]) || isSymbol(matrix[i][j+1]) {
						nextToSymbol = true
					}
				} else if i == 0 && j == len(matrix[0])-1 {
					if isSymbol(matrix[i+1][j]) || isSymbol(matrix[i+1][j-1]) || isSymbol(matrix[i][j-1]) {
						nextToSymbol = true
					}
				} else if i == len(matrix)-1 && j == 0 {
					if isSymbol(matrix[i-1][j]) || isSymbol(matrix[i-1][j+1]) || isSymbol(matrix[i][j+1]) {
						nextToSymbol = true
					}
				} else if i == len(matrix)-1 && j == len(matrix[0])-1 {
					if isSymbol(matrix[i-1][j]) || isSymbol(matrix[i-1][j-1]) || isSymbol(matrix[i][j-1]) {
						nextToSymbol = true
					}
				} else if i == 0 {
					if isSymbol(matrix[i+1][j+1]) || isSymbol(matrix[i+1][j]) || isSymbol(matrix[i+1][j-1]) || isSymbol(matrix[i][j+1]) || isSymbol(matrix[i][j-1]) {
						nextToSymbol = true
					}
				} else if j == 0 {
					if isSymbol(matrix[i+1][j+1]) || isSymbol(matrix[i][j+1]) || isSymbol(matrix[i-1][j+1]) || isSymbol(matrix[i+1][j]) || isSymbol(matrix[i-1][j]) {
						nextToSymbol = true
					}
				} else if i == len(matrix)-1 {
					if isSymbol(matrix[i-1][j+1]) || isSymbol(matrix[i-1][j]) || isSymbol(matrix[i-1][j-1]) || isSymbol(matrix[i][j+1]) || isSymbol(matrix[i][j-1]) {
						nextToSymbol = true
					}
				} else if j == len(matrix[0])-1 {
					if isSymbol(matrix[i+1][j-1]) || isSymbol(matrix[i][j-1]) || isSymbol(matrix[i-1][j-1]) || isSymbol(matrix[i+1][j]) || isSymbol(matrix[i-1][j]) {
						nextToSymbol = true
					}
				} else {
					if isSymbol(matrix[i+1][j]) || isSymbol(matrix[i+1][j+1]) || isSymbol(matrix[i+1][j-1]) || isSymbol(matrix[i][j+1]) || isSymbol(matrix[i][j-1]) || isSymbol(matrix[i-1][j+1]) || isSymbol(matrix[i-1][j]) || isSymbol(matrix[i-1][j-1]) {
						nextToSymbol = true
					}
				}
			} else {
				if nextToSymbol {
					res *= currentNumber
					conteur++
				}
				currentNumber = 0
				nextToSymbol = false
			}

		}
		if nextToSymbol {
			res *= currentNumber
			conteur++
		}
		currentNumber = 0
		nextToSymbol = false
	}
	if conteur == 2 {
		return res
	} else {
		return 0
	}
}

func isSymbol(ch string) bool {
	if ch == "." || ch == "0" || ch == "1" || ch == "2" || ch == "3" || ch == "4" || ch == "5" || ch == "6" || ch == "7" || ch == "8" || ch == "9" {
		return false
	}
	return true
}
