package main

import "fmt"

func main() {
	var s, a, b, c, res int
	fmt.Scan(&s)
	a = s / 100
	b = s / 10 % 10
	c = s % 10
	res = a + b + c
	fmt.Println(res)
}
