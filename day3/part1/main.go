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
		sum := sumWrongBackpackContents(scanner)

		fmt.Println("SUM", filepath, sum)
	}
}

func sumWrongBackpackContents(scanner *bufio.Scanner) int {
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		middle := len(line) / 2
		compartmentOne, compartmentTwo := line[:middle], line[middle:]

		visited := [53]bool{false}
		for _, itemOne := range compartmentOne {
			priority := itemPriority(itemOne)

			if visited[priority] {
				continue
			}
			visited[priority] = true

			for _, itemTwo := range compartmentTwo {
				if itemOne != itemTwo {
					continue
				}
				sum += priority
				break
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
