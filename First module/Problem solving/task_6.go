package main

import "fmt"

func main() {
	var a, b float32
	fmt.Scan(&a, &b)
	var c float32 = (a + b) / 2.0
	fmt.Println(c)
}
