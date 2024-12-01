package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var list1 []int
	var list2 []int

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	// For each line in input, create separate lists.
	for scanner.Scan() {
		line := scanner.Text()
		listEntries := strings.Split(line, "   ")
		listEntry1, _ := strconv.Atoi(listEntries[0])
		listEntry2, _ := strconv.Atoi(listEntries[1])
		list1 = append(list1, listEntry1)
		list2 = append(list2, listEntry2)
	}

	partOne(list1, list2)
	partTwo(list1, list2)
}

// Find the answer for part 1!
func partOne(list1 []int, list2 []int) {
	sortSlice(list1)
	sortSlice(list2)

	totalDistance := 0

	for i := range list1 {
		distance := math.Abs(float64(list1[i]) - float64(list2[i]))
		totalDistance = totalDistance + int(distance)
	}

	fmt.Printf("The answer to day 1, part 1 is %d!\n", totalDistance)
}

// Sort int arrays in ascending order.
func sortSlice(numberList []int) {
	sort.Slice(numberList, func(x, y int) bool {
		return numberList[x] < numberList[y]
	})
}

// Find the answer for part 2!
func partTwo(list1 []int, list2 []int) {
	freqMap := createFreqMap(list2)
	simScore := 0

	for _, number := range list1 {
		simScore = simScore + (int(number) * freqMap[number])
	}

	fmt.Printf("The answer to day 1, part 2 is %d!", simScore)
}

// Make a frequency map for all numbers in a list.
func createFreqMap(numberList []int) map[int]int {
	freqMap := make(map[int]int)

	for _, number := range numberList {
		freqMap[number] = freqMap[number] + 1
	}

	return freqMap
}
