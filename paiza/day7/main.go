package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	// [o] 総量 [1] 最初に売った量 [2]次に売った量
	firstLine := []float64{}
	for _, v := range strings.Split(sc.Text(), " ") {
		vi, _ := strconv.ParseFloat(v, 64)
		firstLine = append(firstLine, vi)
	}

	firstSold := firstLine[1] / 100.0
	secondSold := firstLine[2] / 100.0
	firstSoldAmount := firstLine[0] - (firstSold * firstLine[0])
	result := firstSoldAmount - (firstSoldAmount * secondSold)

	fmt.Println(result)
}