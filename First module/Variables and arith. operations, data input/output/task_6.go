package main

import "fmt"

func main() {
	var d int
	fmt.Scan(&d)

	hours := d * 2 / 60
	minutes := d * 2 % 60
	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}
