package main

import "fmt"

func main() {
	var s, a, b, c int
	fmt.Scan(&s)
	a = s / 100
	b = s / 10 % 10
	c = s % 10
	fmt.Printf("%d%d%d", c, b, a)
}
