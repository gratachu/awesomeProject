package main

import "fmt"

func main()  {
	var N int
	fmt.Scan(&N)

	arr := []int{}

	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)

		arr = append(arr, num)
	}

	resultArr := Uniq(arr)
	fmt.Println(len(resultArr))
}

func Uniq(words []int) []int {
	m := make(map[int]bool)
	uniq := [] int{}

	for _, ele := range words {
		if !m[ele] {
			m[ele] = true
			uniq = append(uniq, ele)
		}
	}

	return uniq
}