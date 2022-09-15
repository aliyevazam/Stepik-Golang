package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64
	fmt.Scan(&s)
	if s <= 0 {
		fmt.Printf("число %.2f не подходит", s)
	} else if s >= 10000 {
		fmt.Printf("%e", s)
	} else {
		fmt.Printf("%.4f", float64(math.Pow(s, float64(2))))
	}
}
