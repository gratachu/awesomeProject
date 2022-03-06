package main

import (
	"fmt"
	"sort"
)

func main()  {
	var N int
	fmt.Scan(&N)

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	// 配列をソートする
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	// get cart first alice second bob

	var Alice, Bob int
	for i := 0; i < N; i++ {
		if i % 2 == 0 {
			Alice += arr[i]
		} else {
			Bob += arr[i]
		}
	}

	fmt.Println(Alice - Bob)
}