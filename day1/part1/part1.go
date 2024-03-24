package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
    if len(os.Args) <= 1 {
        fmt.Println("ERROR", "\"please pass the file to be processed as command line argument\"")
    }

    for _, filepath := range os.Args[1:] {
        file, err := os.Open(filepath)
        if err != nil {
            fmt.Println("ERROR", "\"could not open file\"", filepath, err)
            continue
        }
        defer file.Close()

        reader := bufio.NewScanner(file)
        sum, err := SumOfMaxElf(reader)
        if err != nil {
            fmt.Println("ERROR", err)
        }

        fmt.Println("RESULT", filepath, sum)
    }
}

func SumOfMaxElf(scanner *bufio.Scanner) (int, error) {
    var sumCurrentElf int = 0
    var maxSumElf int = 0

    for lineNumber := 1; scanner.Scan() ; lineNumber++ {
        line := scanner.Text()

        if line != "" {
            number, err := strconv.Atoi(line)
            if err != nil {
                err := errors.New(fmt.Sprintf("could not parse integer from line %d with content %s with inner error %w", lineNumber, line, err))
                return maxSumElf, err
            }
            sumCurrentElf += number
            continue
        }

        if sumCurrentElf > maxSumElf {
            maxSumElf = sumCurrentElf
        }
        sumCurrentElf = 0
    }

    if sumCurrentElf > maxSumElf {
        maxSumElf = sumCurrentElf
    }

    return maxSumElf, nil
}
