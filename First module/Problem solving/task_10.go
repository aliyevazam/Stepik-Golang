package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b)
	for i := b; i >= a; i-- {
		if i%7 == 0 {
			fmt.Println(i)
			c++
			break
		}
	}
	if c == 0 {
		fmt.Println("NO")
	}
}
