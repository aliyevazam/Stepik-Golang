package main

import "fmt"

func main() {
	var n, min, s int
	fmt.Scan(&n)
	ar := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}
	min = ar[0]
	for i := 0; i < n; i++ {
		if ar[i] < min {
			min = ar[i]
		}
	}
	for i := 0; i < n; i++ {
		if ar[i] == min {
			s++
		}
	}
	fmt.Println(s)
}
