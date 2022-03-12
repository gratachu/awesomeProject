package main

import "fmt"

func main()  {
	board := make(map[int]string)

	for i := 0; i < 5; i++ {
		var l string
		fmt.Scan(&l)

		board[i] = l
	}

	win := JudgeWinner(board)

	fmt.Println(win)
}

func JudgeWinner(board map[int]string) string {
	for _, b := range board {
		switch {
		case b == "OOOOO":
			return "O"
		case b == "XXXXX":
			return "X"
		}
	}

	return "D"
}
