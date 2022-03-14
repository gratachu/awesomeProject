package main

import (
	"fmt"
	"strings"
)

func main()  {
	var str string
	fmt.Scan(&str)

	str = strings.ReplaceAll(str, "a", "")
	str = strings.ReplaceAll(str, "i", "")
	str = strings.ReplaceAll(str, "u", "")
	str = strings.ReplaceAll(str, "e", "")
	str = strings.ReplaceAll(str, "o", "")
	str = strings.ReplaceAll(str, "A", "")
	str = strings.ReplaceAll(str, "I", "")
	str = strings.ReplaceAll(str, "U", "")
	str = strings.ReplaceAll(str, "E", "")
	str = strings.ReplaceAll(str, "O", "")

	fmt.Println(str)
}