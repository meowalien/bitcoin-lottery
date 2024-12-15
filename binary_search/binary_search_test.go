package binary_search

import (
	"os"
	"testing"
)

func createTestFile(t *testing.T, lines []string) *os.File {
	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			t.Fatalf("Failed to write to temp file: %v", err)
		}
	}

	err = file.Close()
	if err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	file, err = os.Open(file.Name())
	if err != nil {
		t.Fatalf("Failed to reopen temp file: %v", err)
	}

	return file
}

func TestBinarySearchOnDisk(t *testing.T) {
	lines := []string{"apple", "banana", "cherry", "date", "fig", "grape"}
	file := createTestFile(t, lines)
	defer os.Remove(file.Name())

	tests := []struct {
		target   string
		expected bool
	}{
		{"banana", true},
		{"fig", true},
		{"kiwi", false},
		{"apple", true},
		{"grape", true},
		{"cherry", true},
		{"date", true},
		{"mango", false},
	}

	for _, test := range tests {
		result, err := BinarySearchOnDisk(file, test.target)
		if err != nil {
			t.Errorf("Error searching for %s: %v", test.target, err)
		}
		if result != test.expected {
			t.Errorf("Expected %v for target %s, got %v", test.expected, test.target, result)
		}
	}
}
