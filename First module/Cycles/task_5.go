package main

import "fmt"

func main() {
	var n, c, d int
	fmt.Scan(&n, &c, &d)
	for kr := 0; kr <= n; {
		if kr%c == 0 && kr%d != 0 {
			fmt.Print(kr)
			break
		}
		kr++
	}
}
