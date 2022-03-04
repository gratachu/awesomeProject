package main

import "fmt"

func main()  {
	var N int
	res := 0
	fmt.Scanf("%d", &N)

	Judge(N, res)
}

func Judge(n, res int) {
	for cake := 0; cake < 101; cake++{
		for donuts := 0; donuts < 101; donuts++ {
			total := 4 * cake + 7 * donuts
			if total == n {
				res++
			}
		}
	}	

	if 0 < res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}