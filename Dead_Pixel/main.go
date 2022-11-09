package main

import "fmt"
import "math"

func solve() {
	var a, b, x, y float64
	fmt.Scan(&a, &b, &x, &y)
	res1 := math.Max(float64(x * b), float64((a - 1 - x) * b))
	res2 := math.Max(float64(a * y), float64(a * (b - 1 - y)))
	res := math.Max(res1, res2)
	fmt.Printf("%d\n", int(res))
}

func main() {
	var t int64
	fmt.Scan(&t)
	for ; t > 0; t-- {
		solve()
	}
}