package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	left := n/100000%10 + n/10000%10 + n/1000%10
	right := n/100%10 + n/10%10 + n%10
	if left == right {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
