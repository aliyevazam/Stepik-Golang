package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a < 10 {
		// a = a
	} else if a < 100 {
		a = a / 10
	} else if a < 1000 {
		a = a / 100
	} else if a < 10000 {
		a = a / 1000
	} else if a == 10000 {
		a = a / 10000
	}

	fmt.Println(a)
}
