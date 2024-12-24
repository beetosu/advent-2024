package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var rulesList []string
	var updateStrings []string

	args := os.Args
	if len(args) > 1 && args[1] == "test" {
		rulesList, updateStrings = testInput()
	} else {
		rulesList, updateStrings = readInput()
	}

	rulesMap := buildRules(rulesList)
	updates := parseUpdateStrings(updateStrings)

	correctMiddlePageCount := 0
	incorrectMiddlePageCount := 0
	fmt.Println(rulesMap)
	for _, update := range updates {
		if updateIsValid(update, rulesMap) {
			middleIdx := int(len(update) / 2)
			middlePage, _ := strconv.Atoi(update[middleIdx])
			correctMiddlePageCount += middlePage
		} else {
			reorderedUpdate := reorderUpdate(update, rulesMap)
		}
	}
	fmt.Println("PART ONE: ", correctMiddlePageCount)
	fmt.Println("PART TWO: ", incorrectMiddlePageCount)
}

func readInput() ([]string, []string) {
	var rulesList []string
	var updates []string
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)
	onRules := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			onRules = false
			continue
		}

		if onRules {
			rulesList = append(rulesList, line)
		} else {
			updates = append(updates, line)
		}
	}
	return rulesList, updates
}

func testInput() ([]string, []string) {
	inputPhases := strings.Split("47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13\n\n75,47,61,53,29\n97,61,53,29,13\n75,29,13\n75,97,47,61,53\n61,13,29\n97,13,75,29,47", "\n\n")
	return strings.Split(inputPhases[0], "\n"), strings.Split(inputPhases[1], "\n")
}

func buildRules(rulesList []string) map[int][]int {
	rulesMap := make(map[int][]int)
	for _, ruleStr := range rulesList {
		ruleList := parseRule(ruleStr)
		existingArr, alreadyExists := rulesMap[ruleList[0]]
		if !alreadyExists {
			rulesMap[ruleList[0]] = []int{ruleList[1]}
		} else {
			rulesMap[ruleList[0]] = append(existingArr, ruleList[1])
		}
	}
	return rulesMap
}

func parseUpdateStrings(updateStrings []string) [][]string {
	var updates [][]string
	for _, updateString := range updateStrings {
		updates = append(updates, strings.Split(updateString, ","))
	}
	return updates
}

func parseRule(ruleStr string) [2]int {
	stringList := strings.Split(ruleStr, "|")
	left, _ := strconv.Atoi(stringList[0])
	right, _ := strconv.Atoi(stringList[1])
	return [2]int{left, right}
}

func updateIsValid(update []string, rulesMap map[int][]int) bool {
	for idx, pageByte := range update {
		pageNum, _ := strconv.Atoi(pageByte)
		badPages := rulesMap[pageNum]
		for _, pastPageStr := range update[:idx] {
			pastPage, _ := strconv.Atoi(pastPageStr)
			if slices.Contains(badPages, pastPage) {
				return false
			}
		}
	}
	return true
}

func reorderUpdate(badUpdate []string, rulesMap map[int][]int) []string {
	reorderedUpdate := []string{badUpdate[0]}
	for _, page := range badUpdate {
		// iterate through the bad update, and when a conflict occurs
		// shift the offending pageNum back until it no longer conflicts.
		fmt.Println(page)
	}
	return reorderedUpdate
}
