package main

import "fmt"

func main() {

	fmt.Println("-----------------------------------------------------")
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>ARRAYS<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	fmt.Println("-----------------------------------------------------")
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>SLICE<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	// Khai báo một mảng có 5 phần tử
	arr := [5]int{1, 2, 3, 4, 5}

	// Tạo một slice từ mảng arr, lấy phần tử từ vị trí 1 đến trước vị trí 4
	slice1 := arr[1:4] // Slice sẽ chứa các phần tử 2, 3, 4

	// In ra giá trị của slice1 và mảng arr
	fmt.Println("Mảng gốc:", arr)         // [1 2 3 4 5]
	fmt.Println("Slice ban đầu:", slice1) // [2 3 4]

	// Thay đổi giá trị của phần tử trong slice1
	slice1[1] = 10 // Thay đổi giá trị tại vị trí 1 (tương ứng phần tử thứ 3 của mảng arr)

	// In ra giá trị sau khi thay đổi
	fmt.Println("Mảng gốc sau thay đổi:", arr) // [1 2 10 4 5]
	fmt.Println("Slice sau thay đổi:", slice1) // [2 10 4]

	// Tạo một slice khác từ mảng arr
	slice2 := arr[2:5] // Slice này sẽ chứa các phần tử 10, 4, 5

	// In ra slice2
	fmt.Println("Slice2:", slice2) // [10 4 5]

	fmt.Println("-----------------------------------------------------")
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>Slice Literals, Slice Defaults <<<<<<<<<<<<<<<<<<<<<<<<<<<")

	fmt.Println("-----------------------------------------------------")
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>ARRAYS<<<<<<<<<<<<<<<<<<<<<<<<<<<")
}
