package main

import (
	"fmt"
	"sort"
	"strings"
)

func main()  {
	var str string
	fmt.Scan(&str)

	strs := make([]string, len(str))
	for i, c := range str {
		strs[i] = string(c)
	}

	sort.Strings(strs)
	fmt.Println(strings.Join(strs, ""))
}