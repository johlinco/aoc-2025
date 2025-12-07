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
	freshIngredients := 0
	var ranges [][2]int
	postBreak := false
	for _, line := range lines {
		//fmt.Printf("the line is %s\n", line)
		if line == "" {
			postBreak = true
			continue
		}
		if !postBreak {
			parts := strings.Split(line, "-")
			start, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Printf("Start is Wrong with %s", parts[0])
			}
			end, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Printf("Start is Wrong with %s", parts[1])
			}
			ranges = append(ranges, [2]int{start, end})
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				fmt.Printf("Ingredient Id %s is wrong", line)
			}
			for _, r := range ranges {
				if num >= r[0] && num <= r[1] {
					freshIngredients++
					break
				}
			}
		}
	}
	return freshIngredients
}

func solvePart2(lines []string) int {
	// TODO: Implement part 2
	return 0
}
