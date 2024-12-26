package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	args := os.Args
	var guardMap []string
	xCount := 0
	if len(args) > 1 && args[1] == "test" {
		guardMap = testInput()
	} else {
		guardMap = readInput()
	}

	guardX, guardY := findGuard(guardMap)
	patrolMap := patrol(guardX, guardY, guardMap)
	for _, patrolLine := range patrolMap {
		fmt.Println(patrolLine)
		xCount += strings.Count(patrolLine, "|") + strings.Count(patrolLine, "+") + strings.Count(patrolLine, "-") + strings.Count(patrolLine, "*")
	}
	fmt.Println("PART ONE:", xCount)

	fmt.Println("PART TWO:", findLoops(guardMap, patrolMap, guardX, guardY))
}

func testInput() []string {
	return strings.Split("....#.....\n.........#\n..........\n..#.......\n.......#..\n..........\n.#..^.....\n........#.\n#.........\n......#...", "\n")
}

func readInput() []string {
	var guardMap []string
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		guardMap = append(guardMap, scanner.Text())
	}
	return guardMap
}

func findGuard(guardMap []string) (x int, y int) {
	for y, mapLine := range guardMap {
		x := strings.Index(mapLine, "^")
		if x != -1 {
			return x, y
		}
	}
	return -1, -1
}

func patrol(guardX int, guardY int, guardMap []string) []string {
	patrolMap := make([]string, len(guardMap))
	copy(patrolMap, guardMap)
	parallelTrailAhead := false
	for {
		guardSprite := string(patrolMap[guardY][guardX])
		xDir, yDir := getGuardDir(guardSprite)
		newGuardX := guardX + xDir
		newGuardY := guardY + yDir

		if guardIsOOB(newGuardX, newGuardY, guardMap) {
			patrolMap[guardY] = patrolMap[guardY][:guardX] + "*" + patrolMap[guardY][guardX+1:]
			return patrolMap
		}

		wasPivoted := false
		if string(patrolMap[newGuardY][newGuardX]) == "#" {
			guardSprite = pivotGuard(guardSprite)
			wasPivoted = true
			xDir, yDir = getGuardDir(guardSprite)
			newGuardX = guardX + xDir
			newGuardY = guardY + yDir
		}

		trailToLeave := getGuardTrail(xDir, wasPivoted)
		trailAhead := string(patrolMap[newGuardY][newGuardX])
		if trailAhead == "+" || trailAhead == trailToLeave {
			patrolMap[newGuardY] = patrolMap[newGuardY][:newGuardX] + "G" + patrolMap[newGuardY][newGuardX+1:]
			printTrail(patrolMap)
			return []string{}
		}
		if parallelTrailAhead {
			trailToLeave = "+"
		}
		parallelTrailAhead = guardTrailsParallel(getGuardTrail(xDir, wasPivoted), trailAhead)

		patrolMap[guardY] = patrolMap[guardY][:guardX] + trailToLeave + patrolMap[guardY][guardX+1:]
		patrolMap[newGuardY] = patrolMap[newGuardY][:newGuardX] + guardSprite + patrolMap[newGuardY][newGuardX+1:]
		guardX = newGuardX
		guardY = newGuardY
	}
}

func getGuardDir(guardSprite string) (xDir int, yDir int) {
	switch guardSprite {
	case "^":
		return 0, -1
	case ">":
		return 1, 0
	case "<":
		return -1, 0
	default:
		return 0, 1
	}
}

func guardIsOOB(guardX int, guardY int, guardMap []string) bool {
	mapWidth := len(guardMap[0])
	mapHeight := len(guardMap)
	return !(0 <= guardX && guardX < mapWidth) || !(0 <= guardY && guardY < mapHeight)
}

func pivotGuard(oldGuardSprite string) string {
	switch oldGuardSprite {
	case "^":
		return ">"
	case ">":
		return "v"
	case "v":
		return "<"
	default:
		return "^"
	}
}

func getGuardTrail(xDir int, wasPivoted bool) string {
	if wasPivoted {
		return "+"
	}
	if xDir != 0 {
		return "-"
	}
	return "|"
}

func guardTrailsParallel(trailA string, trailB string) bool {
	return (trailA == "-" && trailB == "|") || (trailA == "|" && trailB == "-")
}

func findLoops(guardMap []string, patrolMap []string, guardX int, guardY int) int {
	loops := 0
	trailCharacters := []string{"|", "-", "+"}
	for y, patrolLine := range patrolMap {
		for x, patrolSpace := range patrolLine {
			if slices.Contains(trailCharacters, string(patrolSpace)) {
				obstructedMap := placeObstruction(guardMap, x, y)
				guardObstructedMap := patrol(guardX, guardY, obstructedMap)

				if len(guardObstructedMap) == 0 {
					loops++
				}
			}
		}
	}
	return loops
}

func placeObstruction(guardMap []string, x int, y int) []string {
	obstructedMap := make([]string, len(guardMap))
	copy(obstructedMap, guardMap)
	obstructedMap[y] = obstructedMap[y][:x] + "#" + obstructedMap[y][x+1:]
	return obstructedMap
}

func printTrail(trail []string) {
	for _, line := range trail {
		fmt.Println(line)
	}
	fmt.Println()
}
