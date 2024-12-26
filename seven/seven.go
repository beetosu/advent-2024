package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	var calibrations map[int][]int
	if len(args) > 1 && args[1] == "test" {
		calibrations = testInput()
	} else {
		calibrations = readInput()
	}

	correctCalibrationSum := 0
	for sum, operands := range calibrations {
		if calibrationCanBeCorrect(sum, operands) {
			correctCalibrationSum += sum
		}
	}
	fmt.Println("PART ONE:", correctCalibrationSum)
}

func testInput() map[int][]int {
	testCalibrations := make(map[int][]int)

	rawTestStrs := strings.Split("190: 10 19\n3267: 81 40 27\n83: 17 5\n156: 15 6\n7290: 6 8 6 15\n161011: 16 10 13\n192: 17 8 14\n21037: 9 7 18 13\n292: 11 6 16 20", "\n")
	for _, calibrationStr := range rawTestStrs {
		result, operands := buildCalibration(calibrationStr)
		testCalibrations[result] = operands
	}
	return testCalibrations
}

func readInput() map[int][]int {
	realCalibrations := make(map[int][]int)
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result, operands := buildCalibration(scanner.Text())
		realCalibrations[result] = operands
	}
	return realCalibrations
}

func buildCalibration(calibrationStr string) (int, []int) {
	pairs := strings.Split(calibrationStr, ": ")
	var operands []int
	result, _ := strconv.Atoi(pairs[0])
	operandStrs := strings.Split(pairs[1], " ")
	for _, operandStr := range operandStrs {
		operand, _ := strconv.Atoi(operandStr)
		operands = append(operands, operand)
	}
	return result, operands
}

func calibrationCanBeCorrect(sum int, operands []int) bool {
	possibleOperators := byte(len(operands) - 1)
	permutationCount := byte(possibleOperators * 2)
	for permutation := range permutationCount {
		value := operands[0]
		for operatorIdx := range possibleOperators {
			operatorFlag := int(permutation>>operatorIdx) % 2
			if operatorFlag == 0 {
				value += operands[operatorIdx+1]
			} else {
				value *= operands[operatorIdx+1]
			}
		}
		if value == sum {
			return true
		}
	}
	return false
}
