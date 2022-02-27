package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 求める座標の構造体
type Coordinate struct {
	X int
	Y int
}

func main()  {
	scanner := bufio.NewScanner(os.Stdin)

	baseData := GetBaseData(scanner)
	if baseData == nil {
		fmt.Printf("baseData is nil")
	}

	// 座標の行数分の文字列を取得する
	texts := GetTexts(scanner, baseData[0])
	// fmt.Println(texts)


	// 求める座標を取得する
	coordinates := GetCoordinates(scanner, baseData[2])
	// fmt.Println(coordinates)

	for _, c := range coordinates {
		str := texts[c.X]
		fmt.Println(str[c.Y:c.Y+1])
	}
}

func GetBaseData(sc *bufio.Scanner) []int {
	sc.Scan()
	str := strings.Split(sc.Text(), " ")
	result := []int{}

	for _, s := range str {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil
		}
		result = append(result, i)
	}

	return result
}

// 求める座標の行数をもらって、座標の構造体のスライスを返す
func GetCoordinates(sc *bufio.Scanner, count int) []Coordinate {
	result := []Coordinate{}

	for i := 0; i < count; i++ {
		sc.Scan()
		line := strings.Split(sc.Text(), " ")
		lineI := ConvStringSliceToIntSlice(line)
		c := Coordinate{}
		c.X = lineI[0]
		c.Y = lineI[1]

		result = append(result, c)
	}

	return result
}

func ConvStringSliceToIntSlice(line []string) []int {
	result := make([]int, len(line))
	for i, v := range line {
		vi, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		result[i] = vi
	}

	return result
}

// 行数もらって、文字列を取得して文字列のスライスで返す
func GetTexts(sc *bufio.Scanner, count int) []string {
	result := []string{}

	for i := 0; i < count; i++ {
		sc.Scan()
		result = append(result, sc.Text())
	}

	return result
}
