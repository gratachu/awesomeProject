package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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

// 文字列のスライス内に指定の文字列が含まれているかどうかをbooleanで返す
func Contain(ss []string, w string) bool {
	for _, v := range ss {
		if v == w {
			return true
		}
	}
	return false
}

// ユニーク処理
func Uniq(words []string) []string {
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

// 一行目に必要な数字、2行目に3行目以降の行数の場合
func GetInfoFirstNum() (winNum string, words []string) {
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