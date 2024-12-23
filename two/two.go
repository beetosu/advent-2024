package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var reports [][]int

	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)

	// For each line in input, create separate lists.
	for scanner.Scan() {
		line := scanner.Text()
		reportStrings := strings.Split(line, " ")
		var reportInts []int
		for _, rs := range reportStrings {
			ri, _ := strconv.Atoi(rs)
			reportInts = append(reportInts, ri)
		}
		reports = append(reports, reportInts)
	}

	partOne(reports)
	partTwo(reports)
}

// Determine how many "safe" reports there are
func partOne(reports [][]int) {
	safeCount := 0
	for _, report := range reports {
		if reportIsSafe(report) {
			safeCount++
		}
	}
	println(safeCount)
}

// Determine how many "safe" reports there are, including reports that would be safe without 1 bad entry
func partTwo(reports [][]int) {
	safeCount := 0
	for _, report := range reports {
		if reportIsSafe(report) {
			safeCount++
			continue
		}
		for idx := range report {
			if reportIsSafe(remove(report, idx)) {
				safeCount++
				break
			}
		}
	}
	fmt.Println(safeCount)
}

// Determine if the report is "safe"
func reportIsSafe(report []int) bool {
	rawSlope := float64(report[1] - report[0])
	if rawSlope == 0 {
		return false
	}
	isIncreasing := rawSlope > 0
	for ridx := range len(report) - 1 {
		levelDiff := report[ridx+1] - report[ridx]
		if isIncreasing && (1 > levelDiff || levelDiff > 3) {
			return false
		}
		if !isIncreasing && (-1 < levelDiff || levelDiff < -3) {
			return false
		}
	}
	return true
}

// Remove just one element from a slice, given a "bad" idx
func remove(slice []int, s int) []int {
	var newSlice []int
	for idx, val := range slice {
		if idx != s {
			newSlice = append(newSlice, val)
		}
	}
	return newSlice
}
