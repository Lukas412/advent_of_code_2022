package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("ERROR", "\"please provide the files to be processed as command line arguments\"")
	}

	for _, filepath := range os.Args[1:] {
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Println("ERROR", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		lineIndex := 0
		for scanner.Scan() {
			lineIndex++
			line := scanner.Text()

			buffer := [14]byte{}
			bufferIndex := 0
			oneRotationComplete := false

			duplicateCount := 0
			letterCounts := [26]uint8{}

			charIndex := 0
			for charIndex = 0; charIndex < len(line); charIndex++ {
				if oneRotationComplete && duplicateCount == 0 {
					break
				}

				char := line[charIndex]

				prevChar := buffer[bufferIndex]
				if prevChar != 0 {
					letterIndex := prevChar - 'a'
					letterCounts[letterIndex]--

					if letterCounts[letterIndex] > 0 {
						duplicateCount--
					}
				}

				buffer[bufferIndex] = char
				letterIndex := char - 'a'
				letterCounts[letterIndex]++

				if letterCounts[letterIndex] > 1 {
					duplicateCount++
				}

				bufferIndex++
				if bufferIndex == len(buffer) {
					bufferIndex = 0
					oneRotationComplete = true
				}
			}

			fmt.Println("RESULT", filepath, lineIndex, charIndex)
		}
	}
}

