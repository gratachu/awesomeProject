package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	// get sequence length and count of query
	var length, queries int
	fmt.Scanf("%d %d", &length, &queries)

	sequence := []int{}
	for i := 0; i < length; i++ {
		var num int
		fmt.Scanf("%d", &num)

		sequence = append(sequence, num)
	}

	queryContent := GetQueryContent(queries)

	for _, q := range queryContent {
		switch {
		case strings.Contains(q, "0"):
			// zeroの処理
			sequence = PushBackDigit(q, sequence)
		case q == "1":
			// 1の場合の処理
			sequence = PopBack(sequence)
		case q == "2":
			// 2の場合の処理
			PrintSequence(sequence)
		}
	}
}

// print sequence
func PrintSequence(seq []int) {
	for i := 0; i < len(seq)-1; i++ {
		fmt.Printf("%d ", seq[i])
	}

	fmt.Printf("%d\n", seq[len(seq)-1])
}

// push back X
func PushBackDigit(str string, seq []int) []int{
	s := strings.Split(str, " ")[1] 
	digit, _ := strconv.Atoi(s)
	seq = append(seq, digit)
	return seq
}

// pop back
func PopBack(seq []int) []int {
	seq = seq[:len(seq)-1]
	return seq
}

// get query content
func GetQueryContent(q int) map[int]string {
	r := map[int]string{}
	sc := bufio.NewScanner(os.Stdin)

	for i := 0; i < q; i++ {
		sc.Scan()
		r[i] = sc.Text()
	}

	return r
}

// func ConvStringSliceToIntSlice(line []string) []int {
// 	result := make([]int, len(line))
// 	for i, v := range line {
// 		vi, err := strconv.Atoi(v)
// 		if err != nil {
// 			panic(err)
// 		}

// 		result[i] = vi
// 	}

// 	return result
// }

// // get maximun digits
// func GetResult(digits []int) int {
// 	sort.Ints(digits)
// 	return digits[2] - digits[0]
// }