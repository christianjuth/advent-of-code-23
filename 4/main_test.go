package main

import (
	"fmt"
	"strings"
	"testing"
	// "strings"
)

func Test(t *testing.T) {
	tests := []struct {
		file      []string
    isPartTwo bool
		expected  int
	}{
    // part one
		{
			file: []string {
        "467..114..",
        "...*......",
        "..35..633.",
        "......#...",
        "617*......",
        ".....+.58.",
        "..592.....",
        "......755.",
        "...$.*....",
        ".664.598..",
      },
      isPartTwo: false,
			expected: 4361,
		},

    // part two
		{
			file: []string {
        "..5.......",
        "...*......",
        "..5.......",
      },
      isPartTwo: true,
			expected: 25,
		},
		{
			file: []string {
        "...5......",
        "...*......",
        "..5.......",
      },
      isPartTwo: true,
			expected: 25,
		},
		{
			file: []string {
        "....5.....",
        "...*......",
        "..5.......",
      },
      isPartTwo: true,
			expected: 25,
		},

		{
			file: []string {
        "..5.......",
        "...*......",
        "...5......",
      },
      isPartTwo: true,
			expected: 25,
		},
		{
			file: []string {
        "...5......",
        "...*......",
        "...5......",
      },
      isPartTwo: true,
			expected: 25,
		},
		{
			file: []string {
        "....5.....",
        "...*......",
        "...5......",
      },
      isPartTwo: true,
			expected: 25,
		},

		{
			file: []string {
        ".50.......",
        "...*......",
        "....5.....",
      },
      isPartTwo: true,
			expected: 250,
		},
		{
			file: []string {
        "...5......",
        "...*......",
        "....5.....",
      },
      isPartTwo: true,
			expected: 25,
		},
		{
			file: []string {
        "..50......",
        "...*......",
        "....5.....",
      },
      isPartTwo: true,
			expected: 250,
		},
		{
			file: []string {
        "....5.....",
        "...*......",
        "....5.....",
      },
      isPartTwo: true,
			expected: 25,
		},

		{
			file: []string {
        "....5.....",
        "...*......",
        ".....50...",
      },
      isPartTwo: true,
			expected: 0,
		},

		{
			file: []string {
        "....5.....",
        "...#......",
        "....5.....",
      },
      isPartTwo: true,
			expected: 0,
		},

		{
			file: []string {
        "..5.5.....",
        "...*......",
        "..........",
      },
      isPartTwo: true,
			expected: 25,
		},
		{
			file: []string {
        "..........",
        "...*......",
        "..5.50....",
      },
      isPartTwo: true,
			expected: 250,
		},
	}

	for i, tt := range tests {
    textName := "Test #" + fmt.Sprint(i)
		t.Run(textName, func(t *testing.T) {
      file := strings.Join(tt.file, "\n")
			result := parse(file, tt.isPartTwo)
			if result != tt.expected {
				t.Errorf("Failed %s: expected '%s', got '%s'", textName, fmt.Sprint(tt.expected), fmt.Sprint(result))
			}
		})
	}
}
