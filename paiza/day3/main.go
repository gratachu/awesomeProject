package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// for produh
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	str := scanner.Text()

	// for local check
	// str := "test"

	horizon := len(str) + 2
	fmt.Println(strings.Repeat("+", horizon))
	fmt.Println("+" + str + "+")
	fmt.Println(strings.Repeat("+", horizon))
}