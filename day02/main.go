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

	exampleLines := strings.Split(strings.TrimSpace(string(example)), ",")
	inputLines := strings.Split(strings.TrimSpace(string(input)), ",")

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
	sumInvalidIDs := 0

	for _, line := range lines {
		startStr := strings.Split(line, "-")[0]
		start, err := strconv.Atoi(startStr)
		if err != nil {
			fmt.Printf("start num %s is invalid\n", startStr)
		}
		endStr := strings.Split(line, "-")[1]
		end, err := strconv.Atoi(endStr)
		if err != nil {
			fmt.Printf("end num %s is invalid\n", endStr)
		}
		//fmt.Printf("start is %d and end is %d\n", start, end)
		for num := start; num <= end; num++ {
			numString := strconv.Itoa(num)
			if firstHalfEqualsSecondHalf(numString) {
				//fmt.Printf("adding %d to sum\n", num)
				sumInvalidIDs += num
			}
		}
	}
	return sumInvalidIDs
}

func solvePart2(lines []string) int {
	sumInvalidIDs := 0

	for _, line := range lines {
		startStr := strings.Split(line, "-")[0]
		start, err := strconv.Atoi(startStr)
		if err != nil {
			fmt.Printf("start num %s is invalid\n", startStr)
		}
		endStr := strings.Split(line, "-")[1]
		end, err := strconv.Atoi(endStr)
		if err != nil {
			fmt.Printf("end num %s is invalid\n", endStr)
		}
		//fmt.Printf("start is %d and end is %d\n", start, end)
		for num := start; num <= end; num++ {
			numString := strconv.Itoa(num)
			if isMadeOfRepeats(numString) {
				//fmt.Printf("adding %d to sum\n", num)
				sumInvalidIDs += num
			}
		}
	}
	return sumInvalidIDs
}

func firstHalfEqualsSecondHalf(string string) bool {
	length := len(string)
	if length%2 != 0 {
		return false
	}

	mid := length / 2

	firstHalf := string[:mid]
	secondHalf := string[mid:]

	return firstHalf == secondHalf
}

func isMadeOfRepeats(string string) bool {
	length := len(string)

	breakPoint := length / 2

	for breakPoint > 0 {
		if length%breakPoint == 0 {
			firstSection := string[0:breakPoint]
			for i := breakPoint; i < length; i += breakPoint {
				if string[i:i+breakPoint] != firstSection {
					break
				}
				if i+breakPoint == length {
					return true
				}
			}
		}
		breakPoint--
	}
	return false
}
