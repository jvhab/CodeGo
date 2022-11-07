package main

import (
	"fmt"
)

func solve() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for idx := range a {
		fmt.Scan(&a[idx])
	}
	if n == 1 && a[0]%2 != 0 {
		fmt.Println(-1)
		return
	}
	for idx := range a {
		if a[idx]%2 == 0 {
			fmt.Println(1)
			fmt.Println(idx+1)
			return
		}
	}
	fmt.Println(2)
	fmt.Printf("%d %d", 1, 2)
	fmt.Println("")
}

func main() {
	var t int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		solve()
	}
}