package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main()  {
	sc := bufio.NewScanner((os.Stdin))
	sc.Scan()
	line := strings.Split(sc.Text(), " ")

	// for test
	// line := []string {"5", "3"}

	rand.Seed(time.Now().Unix())

	addictionCount, _ := strconv.Atoi(line[0])
	substractionCount, _ := strconv.Atoi(line[1])

	CreateAddiction(addictionCount)
	CreateSubstraction(substractionCount)
}

func CreateAddiction(count int) {
	for i := 0; i < count; i++ {
		firstNumber := rand.Intn(90)
		secondNumber := firstNumber + rand.Intn(9)
	
		// fmt.Printf("%d + %d =\n", firstNumber, secondNumber)
		fmt.Println(firstNumber, "+", secondNumber, "=")
	}
}

func CreateSubstraction(count int) {
	for i := 0; i < count; i++ {
		firstNumber := rand.Intn(90)
		secondNumber := firstNumber + rand.Intn(9)
	
		// fmt.Printf("%d - %d =\n", secondNumber, firstNumber)
		fmt.Println(secondNumber, "-", firstNumber, "=")
	}
}

