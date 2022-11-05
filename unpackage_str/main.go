package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(Solve(s))
}

func Solve(s string) string {
	var res string
	var dur string
	temp := []rune(s)
	for _, value := range temp {
		if unicode.IsDigit(value) {
			counts, _ := strconv.Atoi(string(value))
			for i := 0; i < counts-1; i++ {
				res += dur
			}
		} else {
			dur = string(value)
			res += string(value)
		}
	}
	return res
}
