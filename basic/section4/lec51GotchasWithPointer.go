package main

import "fmt"

func updateSlice(s []string) {
	s[0] = "Bye"
}

func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	fmt.Println(mySlice) // Kết quả: [Bye There How Are You]
}
