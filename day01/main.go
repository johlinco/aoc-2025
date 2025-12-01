package main

import (
	"fmt"
	"os"
	"strconv"
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
	dial := 50
	code := 0

	for _, line := range lines {
		direction := []rune(line)[0]
		number, _ := strconv.Atoi(line[1:])

		switch direction {
		case 'L':
			dial -= number
			for dial < 0 {
				dial += 100
			}
		case 'R':
			dial += number
			for dial > 99 {
				dial -= 100
			}
		}
		if dial == 0 {
			code++
		}
	}

	return code
}

func solvePart2(lines []string) int {
	dial := 50
	code := 0

	for _, line := range lines {
		direction := []rune(line)[0]
		number, _ := strconv.Atoi(line[1:])

		switch direction {
		case 'L':
			if dial == 0 {
				code += number / 100
			} else if dial <= number {
				code += 1 + (number-dial)/100
			}
			dial = ((dial-number)%100 + 100) % 100
		case 'R':
			if dial == 0 {
				code += number / 100
			} else if dial+number >= 100 {
				code += 1 + (dial+number-100)/100
			}
			dial = (dial + number) % 100
		}
	}

	return code
}
