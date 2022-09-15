package main

import "fmt"

func main() {
	array := [5]int{}
	var a, max int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	max = array[0]
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if array[i] > array[j] && max < array[i] {
				max = array[i]
			} else if array[j] > array[i] && max < array[j] {
				max = array[j]
			} else {
				max = max
			}
		}
	}
	fmt.Println(max)
}
