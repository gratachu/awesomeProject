package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main()  {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	Check(n)
}

func Check(n int) {
	for a := 1; a < 10; a++ {
		for b := 1; b < 10; b++{
			if a * b == n {
				fmt.Println("Yes")
				return
			}
		}
	}

	fmt.Println("No")
}