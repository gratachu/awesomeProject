package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	str := scanner.Text()

	horizon := len(str) + 2
	fmt.Println(strings.Repeat("+", horizon))
	fmt.Println("+" + str + "+")
	fmt.Println(strings.Repeat("+", horizon))
}