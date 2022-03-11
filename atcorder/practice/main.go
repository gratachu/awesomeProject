package main

import "fmt"

func main()  {
	var N, Y int
	fmt.Scan(&N, &Y)

	res10000c := -1
	res5000c := -1
	res1000c := -1

	for a := 0; a <= N; a++ {
		for b := 0; b + a <= N; b++{
			c := N - a - b
			total := a * 10000 + b * 5000 + c * 1000
			if total == Y {
				res10000c = a
				res5000c = b
				res1000c = c
			}
		}
	}

	fmt.Println(res10000c, res5000c, res1000c)
}