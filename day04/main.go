package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read example and input
	example, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	exampleLines := strings.Split(strings.TrimSpace(string(example)), "\n")
	inputLines := strings.Split(strings.TrimSpace(string(input)), "\n")

	// Part 1 Example
	part1Example := solvePart1(exampleLines)
	fmt.Printf("Part 1 Example: %d\n", part1Example)

	// Part 1 Input
	part1Input := solvePart1(inputLines)
	fmt.Printf("Part 1 Input: %d\n", part1Input)

	// Part 2 Example
	part2Example := solvePart2(exampleLines)
	fmt.Printf("Part 2 Example: %d\n", part2Example)

	// Part 2 Input
	part2Input := solvePart2(inputLines)
	fmt.Printf("Part 2 Input: %d\n", part2Input)
}

func solvePart1(lines []string) int {
	accessibleRolls := 0
	grid := parseGrid(lines)
	for i := range grid {
		for j := range grid[i] {
			if forkliftAccessible(grid, i, j) {

				accessibleRolls++
			}
		}
	}
	return accessibleRolls
}

func solvePart2(lines []string) int {
	removedRolls := 0
	grid := parseGrid(lines)

	rollsFound := 1
	rollsToRemove := [][]int{}
	for rollsFound != 0 {
		rollsFound = 0

		for i := range grid {
			for j := range grid[i] {
				if forkliftAccessible(grid, i, j) {
					removedRolls++
					rollsFound++
					rollsToRemove = append(rollsToRemove, []int{i, j})
				}
			}
		}

		for len(rollsToRemove) > 0 {
			col := rollsToRemove[len(rollsToRemove)-1][0]
			row := rollsToRemove[len(rollsToRemove)-1][1]
			grid[row][col] = "."
			rollsToRemove = rollsToRemove[:len(rollsToRemove)-1]
		}
	}

	return removedRolls
}

func parseGrid(lines []string) [][]string {
	grid := [][]string{}

	for _, line := range lines {
		row := []string{}
		for _, char := range line {
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}

	return grid
}

func forkliftAccessible(grid [][]string, posX int, posY int) bool {
	if grid[posY][posX] != "@" {
		return false
	}
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	surroundingRolls := 0
	for _, dir := range dirs {
		checkX := posX + dir[0]
		checkY := posY + dir[1]
		if checkX >= 0 && checkX < len(grid[0]) && checkY >= 0 && checkY < len(grid) && grid[checkY][checkX] == "@" {
			surroundingRolls++
		}
	}
	return surroundingRolls < 4
}
