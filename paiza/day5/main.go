package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	firstLine := strings.Split(scanner.Text(), " ") 
	scanner.Scan()
	secondLine := strings.Split(scanner.Text(), " ")

	// for test
	// firstLine := []string {"10", "2"}
	// secondLine := []string {"6", "2", "0", "7", "1", "3", "5", "3", "2", "6"}
	log := ConvStringSliceToIntSlice(secondLine)

	allPeriod, _ := strconv.Atoi(firstLine[0])
	campPeriod, _ := strconv.Atoi(firstLine[1])

	var campStartDay int // キャンペーン開始日
	average := float32(0) // 平均来場者数
	var campExecCandidateLine []float32 // キャンペーン候補日数

	for i := 0; i <= allPeriod - campPeriod; i++ {
		// まずは全部足して平均を出す
		var totalVisitCount int
		for _, v := range log[i:i+campPeriod] {
			totalVisitCount += v
		}

		periodAverage := float32(totalVisitCount) / float32(campPeriod)
		campExecCandidateLine = append(campExecCandidateLine, periodAverage)

		if average < periodAverage {
			campStartDay = i+1
			average = periodAverage
		}
	}

	fmt.Println(GetMaxCount(campExecCandidateLine), campStartDay)
}

// stringのスライスをintのスライスに置き換える処理
func ConvStringSliceToIntSlice(line []string) []int {
	var result []int
	for _, v := range line {
		vi, _ := strconv.Atoi(v)

		result = append(result, vi)
	}

	return result
}

// スライスから最大値を取得して、個数をカウントする
func GetMaxCount(averages []float32) int {
	var max float32
	for _, v := range averages {
		if max < v {
			max = v
		}
	}

	var count int
	for _, vv := range averages {
		if vv == max {
			count++
		}
	}
	return count
}