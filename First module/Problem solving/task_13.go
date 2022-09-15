package main

import . "fmt"

func main() {
	var n, count, pre int
	Scan(&n)
	a1 := 0
	a2 := 1
	count = -1
	pre = 0
	for n >= a1 {
		a1, a2 = a2, a1+a2
		if a1 == n {
			pre = 1
		}
		count++
	}
	if pre == 1 {
		Print(count)
	} else {
		Print(-1)
	}

}
