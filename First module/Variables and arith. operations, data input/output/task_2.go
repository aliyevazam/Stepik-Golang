package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	a = a * a
	b = b * b
	c := a + b
	fmt.Println(c)
}
