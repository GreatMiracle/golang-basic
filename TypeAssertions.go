package main

import "fmt"

func main() {
	var i interface{} = "Hello, Go!"
	s, ok := i.(string)
	if ok {
		fmt.Printf("Giá trị của s: %s (kiểu string)\n", s)
	} else {
		fmt.Println("Không thể chuyển đổi i thành string")
	}

	// Type assertion thất bại (sẽ không gây panic)
	f, ok := i.(float64)
	if ok {
		fmt.Printf("Giá trị của f: %f (kiểu float64)\n", f)
	} else {
		fmt.Println("Không thể chuyển đổi i thành float64")
	}

	// Type assertion không kiểm tra (gây panic nếu sai kiểu)
	// f2 := i.(float64)  // Dòng này sẽ gây panic nếu không comment lại
	// fmt.Println(f2)
}
