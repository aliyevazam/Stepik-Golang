package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		if a[i] > 0 {
			s++
		}
	}
	fmt.Println(s)
}
