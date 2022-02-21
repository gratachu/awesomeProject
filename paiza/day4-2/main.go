package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	words := GetInfo()
	CallStatus(words)
}

// 最初が行数, 2行目以降の文字列を取得してスライスで返す
func GetInfo() (words []string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	firstStr := scanner.Text()

	count, err := strconv.Atoi(firstStr)
	if err != nil {
		panic(err)
	}

	n := -1

	for scanner.Scan() {
		str := scanner.Text()
		words = append(words, strings.TrimSpace(str))
		n++

		if n >= count {
			break
		}
	}

	return
}

// 文字列のスライスを渡して、期待する文字列を出力する処理
func CallStatus(words []string)  {
	ballCount := 0
	strikeCount := 0

	for _, w := range words {
		switch {
		// ボールカウントが3未満でボールをもらったとき
		case ballCount < 3 && w == "ball":
			fmt.Println("ball!")
			ballCount++
		// ボールカウントが3でボールをもらったとき
		case ballCount == 3 && w == "ball":
			fmt.Println("fourball!")
		// strikeカウントが2未満でストライクをもらう
		case strikeCount < 2 && w == "strike":
			fmt.Println("strike!")
			strikeCount++
		// strikeカウントが2でストライクをもらう
		case strikeCount == 2 && w == "strike":
			fmt.Println("out!")
		}
	}
}