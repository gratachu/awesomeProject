package main

import (
	"fmt"
)

func main()  {
	// N, A, Bを取得
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	var ans int
	for i := 1; i <= n; i++ {
		v := i
		s := 0

		for  v > 0 {
			s += v % 10
			v /= 10
		}
		if s >= a && s <= b {
			ans += i
		}
	}

	fmt.Println(ans)
}