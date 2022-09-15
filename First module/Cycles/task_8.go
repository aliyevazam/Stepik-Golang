package main

import "fmt"

func main() {
	var x int
	var y int
	var k int
	fmt.Scan(&x)
	fmt.Scan(&y)
	for j := 1000; x/j >= 0; j = (j / 10) {
		if x/j == 0 {
			continue
		}
		k = (x / j) % 10
		for i := 1000; y/i >= 0; i = (i / 10) {
			if y/i == 0 {
				continue
			}
			if k == (y/i)%10 {
				fmt.Print(k, " ")
			}
			if i == 1 {
				break
			}
		}
		if j == 1 {
			break
		}
	}
}
