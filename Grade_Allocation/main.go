package main

import "fmt"

func solve() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, n)
	sums := 0
	for idx := range a {
		fmt.Scan(&a[idx])
		sums += a[idx]
	}
	if sums > m {
		fmt.Println(m)
	} else {
		fmt.Println(sums)
	}

}

func main() {
	var t int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		solve()
	}
}