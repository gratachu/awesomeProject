package main

import (
	"fmt"
	"strings"
)

func main()  {
	sample := "red green blue blue green blue"

	ss := strings.Split(sample, " ")
	counter := make(map[string]int)

	// 単語の羅列の中に何個入っているか確認
	for _, s := range ss {
		count := strings.Count(sample, s)
		counter[s] = count
	}

	makeResult(uniq(ss), counter)

}

func makeResult(uniqStrings []string, counter map[string]int) {
	for _, k := range uniqStrings {
		fmt.Printf("%s %d\n" ,k, counter[k])
	}
}
func uniq(ss []string) []string {
	uniqStrings := []string{}
	m := make(map[string]bool)

	for _, e := range ss {
		if !m[e] {
			m[e] = true
			uniqStrings = append(uniqStrings, e)
		}
	}

	return uniqStrings
}