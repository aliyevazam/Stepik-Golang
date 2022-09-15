package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	i := 1
	for i <= n {
		fmt.Printf("%d ", i)
		i = i * 2
	}
}
