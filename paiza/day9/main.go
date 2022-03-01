package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str := sc.Text()
	digits := ConvStr2Slice(str)

	var count int
	
	// attention to infinity loop
	for i := 0; i < 10; {
		if !JudgeDivideAllNumber(digits) {
			break
		}

		digits = DivideAllDigits(digits)
		count++
	}

	fmt.Println(count)
}

// convert string to integer slice
func ConvStr2Slice(s string) []int {
	ss := strings.Split(s, " ")
	r := []int{}

	for _, v := range ss {
		i, _ := strconv.Atoi(v)
		r = append(r, i)
	}

	return r
}

// judge all number of slice can divide
func JudgeDivideAllNumber(digits []int) bool{
	var existOdd bool

	for _, d := range digits {
		if d % 2 != 0 {
			existOdd = false
			return existOdd
		}
	}

	return true
}

func DivideAllDigits(digits []int) []int {
	r := []int{}
	for _, d := range digits {
		i := d / 2
		r = append(r, i)
	}

	return r
}