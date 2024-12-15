package binary_search

import (
	"os"
	"testing"
)

func TestBinarySearchOnDisk1(t *testing.T) {
	//lines := []string{"apple", "banana", "cherry", "date", "fig", "grape"}
	//file := createTestFile(t, lines)
	//defer os.Remove(file.Name())

	// open the file address_after_clean_only_address_sorted.txt
	file, err := os.Open("../address_after_clean_only_address_sorted.txt")
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}

	tests := []struct {
		target   string
		expected bool
	}{
		{"17oscJf4PFTnBkJ9E4xjrT1tNEarLA9jiz", true},
		{"17oscJf4PFTnBkJ9E4xjrT1tNEarLA9jir", false},
		//{"fig", true},
		//{"kiwi", false},
		//{"apple", true},
		//{"grape", true},
		//{"cherry", true},
		//{"date", true},
		//{"mango", false},
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
