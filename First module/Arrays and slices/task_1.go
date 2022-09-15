workArray := [10]uint8{}
var a, b uint8

for i := 0; i < 10; i++ {
	fmt.Scan(&workArray[i])
}
for i := 0; i < 3; i++ {
	fmt.Scan(&a, &b)
	workArray[a], workArray[b] = workArray[b], workArray[a]
}
for _, e := range workArray {
	fmt.Print(e, " ")
}

