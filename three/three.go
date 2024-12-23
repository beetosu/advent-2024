package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fileBytes, _ := os.ReadFile("input")
	fileStr := string(fileBytes)
	partOne(fileStr)
	fmt.Print("\n")
	partTwo(fileStr)
}

func partOne(fileStr string) {
	totalMul := 0
	r := regexp.MustCompile(`mul\([0-9]*,[0-9]*\)`)
	matches := r.FindAllString(fileStr, -1)
	for _, mul := range matches {
		totalMul += evalMul(mul)
	}
	fmt.Print(totalMul)
}

func partTwo(fileStr string) {
	totalMul := 0
	r := regexp.MustCompile(`(mul\([0-9]*,[0-9]*\)|do\(\)|don't\(\))`)
	matches := r.FindAllString(fileStr, -1)
	shouldEval := true
	for _, match := range matches {
		switch match {
		case "do()":
			shouldEval = true
		case "don't()":
			shouldEval = false
		default:
			if shouldEval {
				totalMul += evalMul(match)
			}
		}
	}
	fmt.Print(totalMul)
}

func evalMul(mul string) int {
	justNumbersString := mul[4 : len(mul)-1]
	numbersArr := strings.Split(justNumbersString, ",")
	left, _ := strconv.Atoi(numbersArr[0])
	right, _ := strconv.Atoi(numbersArr[1])
	return left * right
}
