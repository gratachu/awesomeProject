package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var x, y, k, n int

	fmt.Scanf("%d %d", &x, &y)
	fmt.Scan(&k)
	fmt.Scan(&n)

	// 基準点
	baseCoordinate := []int{x, y}

	// 与えられる座標群
	coordinates := [][]int{}

	// 地価の合計
	var landPriceSum int

	// a: x座標, b: y座標, c: 地価
	for i := 0; i < n; i++ {
		var a, b, c int
		fmt.Scanf("%d %d %d", &a, &b, &c)

		s := []int{a, b, c}
		coordinates = append(coordinates, s)
	}

	for index, c := range coordinates {
		abs := math.Abs(float64(c[0])-float64(baseCoordinate[0])) + math.Abs(float64(c[1])-float64(baseCoordinate[1]))
		coordinates[index] = append(coordinates[index], int(abs))
	}

	// 0: x, 1: y, 2: landPlace, 3: abs
	sort.Slice(coordinates, func(i, j int) bool {
		return coordinates[i][3] <= coordinates[j][3]
	})

	useCoordinates := coordinates[0:k]

	for _, u := range useCoordinates {
		landPriceSum += u[2]
	}

	fmt.Println(landPriceSum / k)
}
