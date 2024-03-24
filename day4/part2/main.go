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

		firstRange, err := assignmentRange(firstAssignment)
		if err != nil {
			return sum, err
		}

		secondRange, err := assignmentRange(secondAssignment)
		if err != nil {
			return sum, err
		}

		firstOverlapsSecond := firstRange.Overlaps(secondRange)
		secondOverlapsFirst := secondRange.Overlaps(firstRange)

		if firstOverlapsSecond || secondOverlapsFirst {
			sum++
		}
	}

	return sum, nil
}

func assignmentRange(assignment string) (JobRange, error) {
	startString, endString, err := Split2(assignment, '-')
	if err != nil {
		return  DefaultRange(), err
	}

	start, err := strconv.Atoi(startString)
	if err != nil {
		return DefaultRange(), err
	}

	end, err := strconv.Atoi(endString)

	return NewRange(start, end), err
}

func Split2(value string, separator rune) (string, string, error) {
	index := strings.IndexRune(value, separator)
	if index == -1 {
		return "", "", fmt.Errorf("Cannot split \"%s\" by char '%c'", value, separator)
	}
	return value[:index], value[index + 1:], nil
}

type JobRange struct {
	start int
	end int
}

func DefaultRange() JobRange {
	return JobRange { start: 0, end: 0 }
}

func NewRange(start int, end int) JobRange {
	return JobRange { start: start, end: end }
}

func (r JobRange) Overlaps(other JobRange) bool {
	if r.start <= other.start && other.start <= r.end {
		return true
	}
	return r.start <= other.end && other.end <= r.end
}

