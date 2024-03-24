package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("ERROR", "\"please provide the files to be processed as command line arguments\"")
	}

	for _, filepath := range os.Args[1:] {
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Println("ERROR", err)
			continue
		}
		defer file.Close()
		
		scanner := bufio.NewScanner(file)

		game, err := ParseGame(scanner)
		if err != nil {
			fmt.Println("ERROR", filepath, "\"error while parsing game\"", err)
			continue
		}

		game.Run()
		result := game.TopRow()

		fmt.Println("RESULT", filepath, result)
	}
}

type Game struct {
	stacks [][]byte
	instructions []Instruction
}

func ParseGame(scanner *bufio.Scanner) (Game, error) {
	game := Game {}
	
	crateLines := []string{}
	stacksCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		stackNumberLine := line[1] == '1'
		if stackNumberLine {
			for _, char := range line {
				if char != ' ' {
					stacksCount++
				}
			}
			continue
		}

		crateLines = append(crateLines, line)
	}

	for index := 0; index < stacksCount; index++ {
		game.stacks = append(game.stacks, []byte{})
	}

	for _, line := range crateLines {
		for index := 0; index < stacksCount; index++ {
			stackIndex := 1 + index * 4
			if stackIndex >= len(line) {
				break
			}

			crateValue := line[stackIndex]
			if crateValue == ' ' {
				continue
			}

			game.stacks[index] = append(game.stacks[index], crateValue)
		}
	}

	for _, stack := range game.stacks {
		slices.Reverse(stack)
	}

	for scanner.Scan() {
		line := scanner.Text()
		
		instruction, err := ParseInstruction(line)
		if err != nil {
			return game, err
		}

		game.instructions = append(game.instructions, instruction)
	}

	return game, nil
}

func (game *Game) Run() {
	for _, instruction := range game.instructions {
		fromIndex := instruction.fromStack - 1
		fromLen := len(game.stacks[fromIndex])
		toIndex := instruction.toStack - 1

		fromStartIndex := fromLen - instruction.amount

		game.stacks[toIndex] = append(game.stacks[toIndex], game.stacks[fromIndex][fromStartIndex:fromLen]...)
		game.stacks[fromIndex] = slices.Delete(game.stacks[fromIndex], fromStartIndex, fromLen)
	}

	game.instructions = slices.Delete(game.instructions, 0, len(game.instructions))
}

func (game Game) TopRow() string {
	result := []byte{}
	for _, stack := range game.stacks {
		lastCrate := stack[len(stack) - 1]
		result = append(result, lastCrate)
	}
	return string(result)
}

type Instruction struct {
	fromStack int
	toStack int
	amount int
}

func ParseInstruction(value string) (Instruction, error) {
	values := strings.SplitN(value, " ", 6)
	
	amount, err := strconv.Atoi(values[1])
	if err != nil {
		return NewInvalidInstruction(), err
	}

	fromStack, err := strconv.Atoi(values[3])
	if err != nil {
		return NewInvalidInstruction(), err
	}

	toStack, err := strconv.Atoi(values[5])
	if err != nil {
		return NewInvalidInstruction(), err
	}

	return Instruction {
		fromStack: fromStack,
		toStack: toStack,
		amount: amount,
	}, nil
}

func NewInvalidInstruction() Instruction {
	return Instruction { fromStack: -1, toStack: -1, amount: -1 }
}

