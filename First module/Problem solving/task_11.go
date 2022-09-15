package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if 5 <= n%10 && n%10 <= 9 || n%10 == 0 || n/10 == 1 {
		fmt.Print(n, " korov")
	} else if n%10 == 1 {
		fmt.Print(n, " korova")
	} else {
		fmt.Print(n, " korovy")
	}
}
