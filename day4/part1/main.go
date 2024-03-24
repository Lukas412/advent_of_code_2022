package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("ERROR", "\"please provide the filepaths to be processed as arguments\"")
	}

	for _, filepath := range os.Args[1:] {
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Println("ERROR", err)
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		sum, err := sumFullyContainedAssignments(scanner)

		if err != nil {
			fmt.Println("ERROR", err)
		}

		fmt.Println("SUM", filepath, sum)
	}
}

func sumFullyContainedAssignments(scanner *bufio.Scanner) (int, error) {
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		firstAssignment, secondAssignment, err := Split2(line, ',')
		if err != nil {
			return sum, err
		}

		firstStart, firstEnd, err := assignmentRange(firstAssignment)
		if err != nil {
			return sum, err
		}

		secondStart, secondEnd, err := assignmentRange(secondAssignment)
		if err != nil {
			return sum, err
		}

		firstContainsSecond := firstStart <= secondStart && secondEnd <= firstEnd
		secondContainsFirst := secondStart <= firstStart && firstEnd <= secondEnd

		if firstContainsSecond || secondContainsFirst {
			sum++
		}
	}

	return sum, nil
}

func assignmentRange(assignment string) (int, int, error) {
	startString, endString, err := Split2(assignment, '-')
	if err != nil {
		return 0, 0, err
	}

	start, err := strconv.Atoi(startString)
	if err != nil {
		return start, 0, err
	}

	end, err := strconv.Atoi(endString)

	return start, end, err
}

func Split2(value string, separator rune) (string, string, error) {
	index := strings.IndexRune(value, separator)
	if index == -1 {
		return "", "", fmt.Errorf("Cannot split \"%s\" by char '%c'", value, separator)
	}
	return value[:index], value[index + 1:], nil
}

