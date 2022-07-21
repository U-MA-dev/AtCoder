package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	isWhiteStart := true
	for h := 0; h < n; h++ {
		isWhiteStart = !isWhiteStart
		for ht := 0; ht < a; ht++ {
			isWhite := isWhiteStart
			for w := 0; w < n; w++ {
				isWhite = !isWhite
				for wt := 0; wt < b; wt++ {
					fmt.Print(getTile(isWhite))
				}
			}
			fmt.Println()
		}
	}
}

func getTile(isWhite bool) string {
	if isWhite {
		return "."
	}
	return "*"
}
