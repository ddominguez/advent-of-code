package main

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "example_1",
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected: 7,
		},
		{
			name:     "example_2",
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: 5,
		},
		{
			name:     "example_3",
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := part1(tt.input)
			if res != tt.expected {
				t.Fatalf("expected %d, got %d", tt.expected, res)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "example_1",
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected: 19,
		},
		{
			name:     "example_2",
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: 23,
		},
		{
			name:     "example_3",
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: 26,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := part2(tt.input)
			if res != tt.expected {
				t.Fatalf("expected %d, got %d", tt.expected, res)
			}
		})
	}
}
