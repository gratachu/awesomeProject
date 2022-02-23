package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	winNum, words := GetInfo()
	CheckWinning(winNum, words)
}

// 一行目に必要な数字、2行目に3行目以降の行数の場合
func GetInfo() (winNum string, words []string) {
	scanner := bufio.NewScanner(os.Stdin)

	// 当選番号
	scanner.Scan()
	str := scanner.Text()
	winNum = str

	n := -1
	var count int
	for scanner.Scan() {
		str = scanner.Text()
		if n == -1 {
			count, _ = strconv.Atoi(strings.TrimSpace(str))
		} else {
			words = append(words, strings.TrimSpace(str))
		}
		n++

		if n >= count {
			break
		}
	}

	return
}

// 当選番号と購入番号の文字列の配列を渡して当たりか外れかを当てる
func CheckWinning(winNum string, words []string)  {
 winNumInt, _ := strconv.Atoi(winNum)

	for _, v := range words {
		num, _ := strconv.Atoi(v)

		switch {
		case num == winNumInt:
			fmt.Println("first")
		case num + 1 == winNumInt || num -1 == winNumInt:
			fmt.Println("adjacent")
		case num % 10000 == winNumInt % 10000:
			fmt.Println("second")
		case num % 1000 == winNumInt % 1000:
			fmt.Println("third")
		default:
			fmt.Println("blank")
		}
	}
}