package main

import (
	"fmt"
	"strings"
)

func main() {
	board := CreateBoard()
	ansL2R := CreateDiagonalLineLeft2Right(board)
	ansR2L := CreateDiagonalLineRight2Left(board)

	if ansL2R == "O" || ansR2L == "X" {
		fmt.Println(ansL2R)
	} else {
		fmt.Println(ansR2L)
	}
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

// func JudgeVerticalLine(board []string) string {
// 	var winner string
// 	for i := 0; i < 5; i++ {
// 		var line string
// 		for _, s := range board {
// 			ss := strings.Split(s, "")
// 			line += ss[i]
// 		}

// 		winner = CheckWinner(line)
// 	}

// 	return winner
// }

// 左から右の斜めのラインを作る
func CreateDiagonalLineLeft2Right(board []string) string {
	var line string
	for i := 0; i < 5; i++ {
		str := board[i]
		ss := strings.Split(str, "")
		line += ss[i]
	}

	ans := CheckWinner(line)
	return ans
}

// 右から左の斜めのラインを作る
func CreateDiagonalLineRight2Left(board []string) string {
	var line string

	for i := 5; i > 0; i-- {
		str := board[i-1]
		ss := strings.Split(str, "")
		line += ss[i-1]
	}

	ans := CheckWinner(line)
	return ans
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
