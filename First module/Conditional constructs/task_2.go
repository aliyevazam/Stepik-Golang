package main

import "fmt"

func main() {
	var s int
	fmt.Scan(&s)
	a := s / 100
	b := s / 10 % 10
	c := s % 10
	if a == b || a == c || b == c || b == a || c == b || c == a {
		fmt.Print("NO")
	} else {
		fmt.Print("YES")
	}
}
