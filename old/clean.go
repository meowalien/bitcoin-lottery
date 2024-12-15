package main

import (
	"bufio"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"log"
	"os"
	"strings"
)

// 主程式
func main() {
	// create or open a file named address_after_clean.txt, if the file exists, it will be overwritten
	address_after_clean_file, err := os.Create("address_after_clean.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer address_after_clean_file.Close()

	// read line by line from file address.txt
	file, err := os.Open("address.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		addrsAndBalence := strings.Split(line, "\t")

		if len(addrsAndBalence) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		address := addrsAndBalence[0]

		var yes bool
		yes, _ = validateP2PKHAddress(address)
		if yes {
			//fmt.Println(address, " ", addrsAndBalence[1])
			_, err = address_after_clean_file.WriteString(line + "\n")
			if err != nil {
				panic(err)
			}
			err = address_after_clean_file.Sync()
			if err != nil {
				panic(err)
			}
		}

		//yes = ValidateP2WPKHAddress(address) != nil
		//if yes {
		//	fmt.Println(address)
		//	address_after_clean_file.WriteString(address + "\n")
		//}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// validateP2PKHAddress checks if a Bitcoin address is a valid P2PKH address.
func validateP2PKHAddress(address string) (bool, error) {
	decoded := base58.Decode(address)
	if len(decoded) != 25 {
		return false, errors.New("invalid address length")
	}

	// Check the prefix (0x00 for mainnet, 0x6f for testnet)
	prefix := decoded[0]
	if prefix != 0x00 && prefix != 0x6f {
		return false, errors.New("invalid prefix")
	}

	// Verify checksum
	payload := decoded[:21]
	checksum := decoded[21:]
	calculatedChecksum := doubleSHA256(payload)[:4]

	if !equalSlice(checksum, calculatedChecksum) {
		return false, errors.New("invalid checksum")
	}

	return true, nil
}

// doubleSHA256 computes SHA256(SHA256(data)).
func doubleSHA256(data []byte) []byte {
	firstHash := sha256Sum(data)
	secondHash := sha256Sum(firstHash)
	return secondHash
}

// sha256Sum computes the SHA-256 hash of the input data.
func sha256Sum(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

// equalSlice compares two byte slices for equality.
func equalSlice(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
