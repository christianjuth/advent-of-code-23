package main

import (
	"testing"
  "fmt"
)

func TestStripTypeScript(t *testing.T) {
	tests := []struct {
		file  string
		expected int
	}{
		{
			file:  "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet",
			expected: 142,
		},
		{
			file:  "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen",
			expected: 281,
		},
    {
      file: "eighthree",
      expected: 83,
    },
	}

	for i, tt := range tests {
    textName := "Test #" + fmt.Sprint(i)
		t.Run(textName, func(t *testing.T) {
			result := parse(tt.file)
			if result != tt.expected {
				t.Errorf("Failed %s: expected '%s', got '%s'", textName, fmt.Sprint(tt.expected), fmt.Sprint(result))
			}
		})
	}
}
