package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
)

type Move int

const (
    Invalid Move = iota - 1
    Rock
    Paper
    Scissors
)

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
var errOwnMove = errors.New("error in own move")

func sumGames(scanner *bufio.Scanner) (int, error) {
    gamesSum := 0

    for scanner.Scan() {
        line := scanner.Text()

        opponentMove, err := charToMove(line[0])
        if err != nil {
            return gamesSum, fmt.Errorf("%w; inner: (%w), line: \"%s\")", errOpponentMove, err, line)
        }
        ownMove, err := charToMove(line[2])
        if err != nil {
            return gamesSum, fmt.Errorf("%w; inner: (%w), line: \"%s\")", errOwnMove, err, line)
        }

        gamesSum += getScoreOfMove(ownMove)
        gamesSum += getScoreOfGameOutcome(ownMove, opponentMove)
    }

    return gamesSum, nil
}

var errCharToMove = errors.New("cannot convert char to move")

func charToMove(char byte) (Move, error) {
    switch char {
    case 'A', 'X':
        return Rock, nil
    case 'B', 'Y':
        return Paper, nil
    case 'C', 'Z':
        return Scissors, nil
    default:
        charValue := rune(char)
        return Invalid, fmt.Errorf("%w; value: '%c'", errCharToMove, charValue)
    }
}

func getScoreOfGameOutcome(ownMove Move, opponentMove Move) int {
    draw := ownMove == opponentMove
    if draw {
        return 3
    }
    won := (ownMove == Rock && opponentMove == Scissors) || (ownMove == Paper && opponentMove == Rock) || (ownMove == Scissors && opponentMove == Paper)
    if won {
        return 6
    }
    return 0
}

func getScoreOfMove(move Move) int {
    switch move {
    case Rock:
        return 1
    case Paper:
        return 2
    case Scissors:
        return 3
    }
    return 0
}
