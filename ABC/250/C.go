package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	xs := make([]int, q)
	for i := range xs {
		fmt.Scan(&xs[i])
	}

	// arr1[i] = i番目にあるBallの値
	arr1 := make([]int, n)
	// arr2[i] = 値iのBallがあるarr1のidx
	arr2 := make([]int, n)
	for i := 0; i < n; i++ {
		arr1[i] = i + 1
		arr2[i] = i + 1
	}

	for _, x := range xs {
		ballIdx := arr2[x-1] - 1
		nBallIdx := ballIdx + 1
		if nBallIdx == n {
			nBallIdx = ballIdx - 1
		}

		nBallVal := arr1[nBallIdx]

		swap(arr1, ballIdx, nBallIdx)
		swap(arr2, x-1, nBallVal-1)
	}

	c := make([]string, n)
	for i := range c {
		c[i] = strconv.Itoa(arr1[i])
	}
	fmt.Println(strings.Join(c, " "))
}

func swap(arr []int, i1, i2 int) {
	swap := arr[i1]
	arr[i1] = arr[i2]
	arr[i2] = swap
}
