package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("-----------------------------------------------------")
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>NIL SLICE<<<<<<<<<<<<<<<<<<<<<<<<<<<")

	var s1 []int
	fmt.Println(s1 == nil) // true
	fmt.Println(len(s1))   // 0
	fmt.Println(cap(s1))

	fmt.Println("-----------------------------------------------------")
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>MAKE SLICE<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	a := make([]int, 5) // Tạo slice với độ dài là 5 (len(a) = 5, cap(a) = 5)
	printSlice("a", a)

	b := make([]int, 0, 5) // len(b) = 0, cap(b) = 5
	printSlice("b", b)

	b = b[:cap(b)] // len(b) = 5, cap(b) = 5
	b = b[1:]      // len(b) = 4, cap(b) = 4
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	//d := c[2:5]
	//printSlice("d", d)

	fmt.Println("-----------------------------------------------------")
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>SLICE OF SLICE<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
