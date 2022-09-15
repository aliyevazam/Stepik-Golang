package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	switch {
	case a > 0:
		fmt.Println("Число положительное")
	case 0 > a:
		fmt.Println("Число отрицательное")
	case 0 == a:
		fmt.Println("Ноль")
	}
}
