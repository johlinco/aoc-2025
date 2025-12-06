package main

import (
	"fmt"
	"math/big"
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
	joltSum := 0
	for _, line := range lines {
		lineJolt := 0
		maxRight := make([]int, len(line))
		currMax := 0
		for i := len(line) - 2; i >= 0; i-- {
			maxRight[i] = max(currMax, int(line[i+1]-'0'))
			currMax = max(currMax, int(line[i+1]-'0'))
		}
		//fmt.Println("max right array is", maxRight)
		for i := 0; i < len(line)-1; i++ {
			num := int(line[i]-'0')*10 + maxRight[i]
			lineJolt = max(lineJolt, num)
		}
		joltSum += lineJolt
	}
	return joltSum
}

func solvePart2(lines []string) *big.Int {
	joltSum := big.NewInt(0)
	for _, line := range lines {
		lineJoltArray := []int{}

		skipsLeft := len(line) - 12
		for _, s := range line {
			num := int(s - '0')
			for len(lineJoltArray) > 0 && lineJoltArray[len(lineJoltArray)-1] < num && skipsLeft > 0 {
				lineJoltArray = lineJoltArray[:len(lineJoltArray)-1]
				skipsLeft--
			}
			lineJoltArray = append(lineJoltArray, num)
		}

		for skipsLeft > 0 && len(lineJoltArray) > 0 {
			lineJoltArray = lineJoltArray[:len(lineJoltArray)-1]
			skipsLeft--
		}

		lineJoltSum := big.NewInt(0)
		ten := big.NewInt(10)
		for i := 0; i < len(lineJoltArray); i++ {
			lineJoltSum.Mul(lineJoltSum, ten)
			lineJoltSum.Add(lineJoltSum, big.NewInt(int64(lineJoltArray[i])))
		}
		joltSum.Add(joltSum, lineJoltSum)
	}
	return joltSum
}
