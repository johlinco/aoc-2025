package main

import (
	"os"
	"strings"
	"testing"
)

func TestPart1Example(t *testing.T) {
	input, err := os.ReadFile("example.txt")
	if err != nil {
		t.Fatalf("failed to read example: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), ",")
	result := solvePart1(lines)

	expected := 1227775554
	if result != expected {
		t.Errorf("Part 1 = %d; want %d", result, expected)
	}
}

func TestPart2Example(t *testing.T) {
	input, err := os.ReadFile("example.txt")
	if err != nil {
		t.Fatalf("failed to read example: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), ",")
	result := solvePart2(lines)

	expected := 4174379265
	if result != expected {
		t.Errorf("Part 2 = %d; want %d", result, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solvePart1(lines)
	}
}

func BenchmarkPart2(b *testing.B) {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solvePart2(lines)
	}
}
