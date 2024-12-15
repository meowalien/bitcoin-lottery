package main

import (
	"bufio"
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/btcsuite/btcutil/base58"
	"os"
	"strings"
)

func countLines(filename string) (uint, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Initialize a line counter
	var lineCount uint = 0

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, nil
}

// 主程式
func main() {
	lineCount, err := countLines("address_after_clean.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(lineCount)

	// read line by line from file address.txt
	file, err := os.Open("address_after_clean.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	filter := bloom.NewWithEstimates(lineCount, 0.0000000001)

	scanner := bufio.NewScanner(file)
	iteration := 0

	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		addrsAndBalence := strings.Split(line, "\t")

		if len(addrsAndBalence) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		address := addrsAndBalence[0]
		addressByteArray := base58.Decode(address)

		filter.Add(addressByteArray)
		if iteration%1000000 == 0 {
			fmt.Println(iteration)
		}

		iteration++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	f, err := os.Create("bloom_filter.bf")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = filter.WriteTo(w)
	if err != nil {
		panic(err)
	}
	err = w.Flush()
	if err != nil {
		panic(err)
	}
}
