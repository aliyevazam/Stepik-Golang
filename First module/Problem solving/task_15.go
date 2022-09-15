package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	for _, e := range a {
		if string(e) == b {
			continue
		}
		fmt.Print(string(e))
	}
}
