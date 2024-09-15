package main

import "fmt"

func doSomething(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("i là số nguyên: %d\n", v)
	case string:
		fmt.Printf("i là chuỗi: %s\n", v)
	case bool:
		fmt.Printf("i là kiểu boolean: %t\n", v)
	default:
		fmt.Printf("Kiểu của i không xác định: %T\n", v)
	}
}

func main() {
	doSomething(10)
	doSomething("Xin chào")
	doSomething(true)
	doSomething(3.14)
}
