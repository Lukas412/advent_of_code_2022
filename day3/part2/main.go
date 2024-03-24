package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("ERROR", "\"please provide the filepath to be processed as command line arguments\"")
		return
	}

	for _, filepath := range os.Args[1:] {
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Println("ERROR", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		sum := sumGroupFlags(scanner)

		fmt.Println("SUM", filepath, sum)
	}
}

func sumGroupFlags(scanner *bufio.Scanner) int {
	sum := 0
	for scanner.Scan() {
		backpack1 := scanner.Text()
		if !scanner.Scan() {
			break
		}
		backpack2 := scanner.Text()
		if !scanner.Scan() {
			break
		}
		backpack3 := scanner.Text()

		visited1 := [53]bool{false}
		for _, item1 := range backpack1 {
			priority := itemPriority(item1)

			if visited1[priority] {
				continue
			}
			visited1[priority] = true

			for _, item2 := range backpack2 {
				if item1 != item2 {
					continue
				}

				wasFound := false
				for _, item3 := range backpack3 {
					if item1 != item3 {
						continue
					}
					sum += priority
					wasFound = true
					break
				}

				if wasFound {
					break
				}
			}
		}
	}
	return sum
}

func itemPriority(char rune) int {
	if 'a' <= char && char <= 'z' {
		return int(char-'a') + 1
	}
	if 'A' <= char && char <= 'Z' {
		return int(char-'A') + 27
	}
	return 0
}
