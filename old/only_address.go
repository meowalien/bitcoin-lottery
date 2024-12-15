package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// 主程式
func main() {
	// create or open a file named address_after_clean.txt, if the file exists, it will be overwritten
	address_after_clean_file, err := os.Create("address_after_clean_only_address.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer address_after_clean_file.Close()

	// read line by line from file address.txt
	file, err := os.Open("address_after_clean.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	iteration := 0
	for scanner.Scan() {
		iteration++
		line := scanner.Text()
		//fmt.Println(line)
		addrsAndBalence := strings.Split(line, "\t")

		if len(addrsAndBalence) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		address := addrsAndBalence[0]
		_, err = address_after_clean_file.WriteString(address + "\n")
		if err != nil {
			panic(err)
		}

		if iteration%1_000_000 == 0 {
			fmt.Println(iteration)
			err = address_after_clean_file.Sync()
			if err != nil {
				panic(err)
			}
		}

		//var yes bool
		//yes, _ = validateP2PKHAddress(address)
		//if yes {
		//	//fmt.Println(address, " ", addrsAndBalence[1])
		//	_, err = address_after_clean_file.WriteString(line + "\n")
		//	if err != nil {
		//		panic(err)
		//	}
		//	err = address_after_clean_file.Sync()
		//	if err != nil {
		//		panic(err)
		//	}
		//}

		//yes = ValidateP2WPKHAddress(address) != nil
		//if yes {
		//	fmt.Println(address)
		//	address_after_clean_file.WriteString(address + "\n")
		//}

	}
	err = address_after_clean_file.Sync()
	if err != nil {
		panic(err)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
