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
		fmt.Println("ERROR", "\"please provide the files to process as command line arguments\"")
		return
	}

	for _, filepath := range os.Args[1:] {
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Println("ERROR", err)
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		sum, _, err := iterDirectory(scanner)
		if err != nil {
			fmt.Println("ERROR", filepath, err)
		}

		fmt.Println("SUM", filepath, sum)
	}
}

func iterDirectory(scanner *bufio.Scanner) (uint, uint, error) {
	var sum uint = 0
	var output uint = 0

	for scanner.Scan() {
		line := scanner.Text()

		const changeDirectory = "$ cd "
		dirName, found := strings.CutPrefix(line, changeDirectory)
		if found {
			if dirName == ".." {
				if sum < 8381165 {
					return output, sum, nil
				}
				if output == 0 || sum < output {
					return sum, sum, nil
				}
				return output, sum, nil
			}

			inner, size, err := iterDirectory(scanner)
			if err != nil {
				return output, sum, err
			}

			if (inner != 0 && inner < output) || output == 0 {
				output = inner
			}

			sum += size
			continue
		}

		const listDirectory = "$ ls"
		const dirPrefix = "dir"
		if line == listDirectory || strings.HasPrefix(line, dirPrefix) {
			continue
		}

		filesize, _, err := split2(line, ' ')
		if err != nil {
			return output, sum, err
		}

		size, err := strconv.Atoi(filesize)
		if err != nil {
			return output, sum, err
		}

		sum += uint(size)
	}

	if sum < 8381165 {
		return output, sum, nil
	}
	if output == 0 || sum < output {
		return sum, sum, nil
	}
	return output, sum, nil
}

func split2(value string, char byte) (string, string, error) {
	index := strings.IndexByte(value, char)
	if index == -1 {
		return "", "", fmt.Errorf("could not split string \"%s\" by char '%c'", value, char)
	}
	return value[:index], value[index+1:], nil
}
