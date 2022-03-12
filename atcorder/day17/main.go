package main

import (
	"fmt"
	"strings"
)

func main() {
	board := CreateBoard()
	winner := JudgeVerticalLine(board)

	fmt.Println(winner)
}

func CreateBoard() []string {
	board := make([]string, 5)
	for i := 0; i < 5; i++ {
		var str string
		fmt.Scan(&str)
		board[i] = str
	}

	return board
}

func JudgeVerticalLine(board []string) string {
	var winner string
	for i := 0; i < 5; i++ {
		var line string
		for _, s := range board {
			ss := strings.Split(s, "")
			line += ss[i]
		}

		winner = CheckWinner(line)
	}

	return winner
}

func CheckWinner(line string) string {
	switch line {
	case "OOOOO":
		return "O"
	case "XXXXX":
		return "X"
	default:
		return "D"
	}
}
