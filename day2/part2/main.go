package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Move int
const (
    InvalidMove Move = iota - 1
    Rock
    Paper
    Scissors
)

func (move Move) Score() int {
    switch move {
    case Rock:
        return 1
    case Paper:
        return 2
    case Scissors:
        return 3
    default:
        return 0
    }
}

type Outcome int
const (
    InvalidOutcome Outcome = iota - 1
    Win
    Draw
    Lose
)

func (outcome Outcome) Score() int {
    switch outcome {
    case Win:
        return 6
    case Draw:
        return 3
    default:
        return 0
    }
}

func main() {
    if len(os.Args) <= 1 {
        fmt.Println("ERROR", "\"give the filepaths to be processed as arguments\"")
        return
    }

    for _, filepath := range os.Args[1:] {
        file, err := os.Open(filepath)
        if err != nil {
            fmt.Println("ERROR", err)
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        sum, err := sumGames(scanner)
        if err != nil {
            fmt.Println("ERROR", filepath, err)
        }

        fmt.Println("SUM", filepath, sum)
    }
}

var errOpponentMove = errors.New("error in opponents move")
var errOutcome = errors.New("error in outcome")

func sumGames(scanner *bufio.Scanner) (int, error) {
    gamesSum := 0

    for scanner.Scan() {
        line := scanner.Text()

        opponentMove, err := charToMove(line[0])
        if err != nil {
            return gamesSum, fmt.Errorf("%w; inner: (%w), line: \"%s\")", errOpponentMove, err, line)
        }
        outcome, err := charToOutcome(line[2])
        if err != nil {
            return gamesSum, fmt.Errorf("%w; inner: (%w), line: \"%s\")", errOutcome, err, line)
        }

        ownMove := moveForOutcome(outcome, opponentMove)

        gamesSum += outcome.Score()
        gamesSum += ownMove.Score()
    }

    return gamesSum, nil
}

var errCharToMove = errors.New("cannot convert char to move")

func charToMove(char byte) (Move, error) {
    switch char {
    case 'A':
        return Rock, nil
    case 'B':
        return Paper, nil
    case 'C':
        return Scissors, nil
    default:
        charValue := rune(char)
        return InvalidMove, fmt.Errorf("%w; value: '%c'", errCharToMove, charValue)
    }
}

var errCharToOutcome = errors.New("cannot convert char to outcome")

func charToOutcome(char byte) (Outcome, error) {
    switch char {
    case 'X':
        return Lose, nil
    case 'Y':
        return Draw, nil
    case 'Z':
        return Win, nil
    default:
        charValue := rune(char)
        return InvalidOutcome, fmt.Errorf("%w; value: '%c'", errCharToOutcome, charValue)
    }
}

func moveForOutcome(outcome Outcome, opponentMove Move) Move {
    if outcome == Draw {
        return opponentMove
    }
    if (outcome == Lose && opponentMove == Paper) || (outcome == Win && opponentMove == Scissors) {
        return Rock
    }
    if (outcome == Lose && opponentMove == Scissors) || (outcome == Win && opponentMove == Rock) {
        return Paper
    }
    return Scissors
}
