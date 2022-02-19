package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	_, words := GetInfo()
	result := uniq(CreateSearchHistory(words))

	for _, v := range result {
		fmt.Println(v)
	}
}

// 最初に行数、改行後文字列の組み合わせ
func GetInfo() (count int, words []string) {
	scanner := bufio.NewScanner(os.Stdin)

	n := -1
	for scanner.Scan() {
		str := scanner.Text()
		if n == -1 {
			count, _ = strconv.Atoi(strings.TrimSpace(str))
		} else {
			words = append(words, strings.TrimSpace(str))
		}
		n += 1

		if n >= count {
			break
		}
	}

	return
}

// 検索履歴を作成する
func CreateSearchHistory(words []string) []string {
	var result []string
	for _, v := range words {
		if contain(words, v) {
			// 重複部分の削除
			result = append([]string{v}, result...)
		} else {
			result = append([]string{v}, result...)
		}
	}

	return result
}

// 文字列が含まれているかどうかをbooleanで返す
func contain(ss []string, w string) bool {
	for _, v := range ss {
		if v == w {
			return true
		}
	}
	return false
}

// ユニーク処理
func uniq(words []string) []string {
	m := make(map[string]bool)
	uniq := [] string{}

	for _, ele := range words {
		if !m[ele] {
			m[ele] = true
			uniq = append(uniq, ele)
		}
	}

	return uniq
}