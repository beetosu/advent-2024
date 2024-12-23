package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	test()
	var crossword []string

	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		crossword = append(crossword, line)
	}

	partOne(crossword)
	partTwo(crossword)
}

func partOne(crossword []string) {
	matchCount := 0
	xMax := len(crossword)
	yMax := len(crossword[0])

	for y, line := range crossword {
		for x, letter := range line {
			if letter != 'X' {
				continue
			}

			for _, mDir := range getValidMDirs(x, y, xMax, yMax) {
				if wordSearch(crossword, "M", [2]int{x, y}, mDir) {
					matchCount++
				}
			}
		}
	}
	fmt.Println("PART ONE: ", matchCount)
}

func partTwo(crossword []string) {
	matchCount := 0
	xMax := len(crossword)
	yMax := len(crossword[0])

	for y, line := range crossword {
		for x, letter := range line {
			if letter != 'A' || !canMakeAnX(x, y, xMax, yMax) {
				continue
			}
			if hasTwoMases(x, y, crossword) {
				matchCount++
			}
		}
	}
	fmt.Println("PART TWO: ", matchCount)
}

func getValidMDirs(x int, y int, xMax int, yMax int) [][2]int {
	var validMDirs [][2]int
	if x+3 < xMax {
		validMDirs = append(validMDirs, [2]int{1, 0})
		if y+3 < yMax {
			validMDirs = append(validMDirs, [2]int{1, 1})
		}
		if y-3 >= 0 {
			validMDirs = append(validMDirs, [2]int{1, -1})
		}
	}
	if x-3 >= 0 {
		validMDirs = append(validMDirs, [2]int{-1, 0})
		if y+3 < yMax {
			validMDirs = append(validMDirs, [2]int{-1, 1})
		}
		if y-3 >= 0 {
			validMDirs = append(validMDirs, [2]int{-1, -1})
		}
	}

	if y+3 < yMax {
		validMDirs = append(validMDirs, [2]int{0, 1})
	}
	if y-3 >= 0 {
		validMDirs = append(validMDirs, [2]int{0, -1})
	}
	return validMDirs
}

func canMakeAnX(x int, y int, xMax int, yMax int) bool {
	return x+1 < xMax && y+1 < yMax && x-1 >= 0 && y-1 >= 0
}

func wordSearch(crossword []string, currentLetter string, originalCoords [2]int, wordDirection [2]int) bool {
	letterCoords := [2]int{originalCoords[0] + wordDirection[0], originalCoords[1] + wordDirection[1]}
	letterToCheck := string(crossword[letterCoords[1]][letterCoords[0]])
	if letterToCheck != currentLetter {
		return false
	}
	switch letterToCheck {
	case "M":
		return wordSearch(crossword, "A", letterCoords, wordDirection)
	case "A":
		return wordSearch(crossword, "S", letterCoords, wordDirection)
	case "S":
		return true
	}
	// this shouldn't ever happen:
	return false
}

func hasTwoMases(x int, y int, crossword []string) bool {
	relativeCoords := [][2]int{{1, 1}, {-1, 1}, {-1, -1}, {1, -1}}
	masCount := 0
	for _, relativeCoord := range relativeCoords {
		possibleMX := x + relativeCoord[0]
		possibleMY := y + relativeCoord[1]
		possibleM := string(crossword[possibleMY][possibleMX])
		if possibleM != "M" {
			continue
		}

		possibleSX := x + relativeCoord[0]*-1
		possibleSY := y + relativeCoord[1]*-1
		possibleS := string(crossword[possibleSY][possibleSX])
		if possibleS != "S" {
			return false
		}
		masCount++
	}
	return masCount == 2
}

func test() {
	rawTest := "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX"
	partOne(strings.Split(rawTest, "\n"))
	partTwo(strings.Split(rawTest, "\n"))
}
