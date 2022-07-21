package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var h, w, r, c, ans int
	fmt.Scan(&h, &w)
	fmt.Scan(&r, &c)

	if r != 1 {
		ans += 1
	}
	if r != h {
		ans += 1
	}
	if c != 1 {
		ans += 1
	}
	if c != w {
		ans += 1
	}

	fmt.Println(ans)
}
