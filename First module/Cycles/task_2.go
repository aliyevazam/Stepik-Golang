package main

import "fmt"

func main() {
	var a, b, sum int
	fmt.Scan(&a, &b)
	for ; a <= b; a++ {
		sum += a
	}
	fmt.Println(sum)
}
