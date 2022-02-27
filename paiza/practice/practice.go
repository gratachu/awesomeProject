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
	firstSlice := GetOneLiner(scanner)

	scanner.Scan()
	secondSlice := GetOneLiner(scanner)

	for _ , v := range secondSlice {
		fmt.Println(firstSlice[v-1])
	}
}

func GetOneLiner(scanner *bufio.Scanner) []int {
	result := []int{}
	scanner.Scan()
	for _, w := range strings.Split(scanner.Text(), " ") {
		i, _ := strconv.Atoi(w)
		result = append(result, i)
	}

	return result
}