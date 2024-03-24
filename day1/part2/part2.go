package main

import (
	"bufio"
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
    var maxSumElves = [3]int{ 0, 0, 0 }

    if !scanner.Scan() {
        return sumCurrentElf, nil
    }

    for lineNumber := 1; ; lineNumber++ {
        line := scanner.Text()
        couldScan := scanner.Scan()

        if line != "" {
            number, err := strconv.Atoi(line)
            if err != nil {
                err := fmt.Errorf("could not parse integer from line %d with content %s with inner error %w", lineNumber, line, err)
                return sum3(maxSumElves), err
            }
            sumCurrentElf += number
            if couldScan {
                continue
            }
        }

        currentSum := sumCurrentElf
        sumCurrentElf = 0

        for index := 0; index < 3; index++ {
            if currentSum > maxSumElves[index] {
                maxSumElves = shiftInsert3(maxSumElves, currentSum, index)
                break
            }
        }

        if !couldScan {
            return sum3(maxSumElves), nil
        }
    }
}

func sum3(array [3]int) int {
    return array[0] + array[1] + array[2]
}

func shiftInsert3(array [3]int, value int, index int) [3]int {
    if index > 2 {
        return array
    }
    if index == 2 {
        array[2] = value
        return array
    }
    array[2] = array[1]
    if index == 1 {
        array[1] = value
        return array
    }
    array[1] = array[0]
    array[0] = value
    return array
}
