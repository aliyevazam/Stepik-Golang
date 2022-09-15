package main

import "fmt"

func main() {
	var n, b, sum int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&b)
		if b >= 10 && b < 100 && b%8 == 0 {
			sum += b
		}
	}
	fmt.Println(sum)
}
