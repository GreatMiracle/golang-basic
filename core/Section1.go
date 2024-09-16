package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var aa, bb int = 3, 4
var i, j int = 1, 2

func main() {
	fmt.Println("kien")

	fmt.Println("Số ngẫu nhiên: ", rand.Intn(100))

	fmt.Println(add(3, 4))                 // Kết quả: 7
	fmt.Println(addParamSampleType(5, 10)) // Kết quả: 15

	a, b := swap("Golang", "Programming")
	fmt.Println(a, b) // Kết quả: Programming Golang

	fmt.Println("-----------------------------------------------------")
	fmt.Println("Named return values")
	total, difference := calc(10, 4)
	fmt.Println("Tổng:", total)
	fmt.Println("Hiệu:", difference)

	fmt.Println("-----------------------------------------------------")
	fmt.Println("Variables")
	var result int = aa + bb
	fmt.Println("Tổng của aa và bb:", result)

	fmt.Println("-----------------------------------------------------")
	fmt.Println("Variables with initializers")
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	fmt.Println("-----------------------------------------------------")
	fmt.Println("Short variable declarations")
	x := 5         // int
	name := "Go"   // string
	active := true // bool
	fmt.Println(x, name, active)

	fmt.Println("-----------------------------------------------------")
	fmt.Println("Basic types")
	var b2 bool = true         // Kiểu bool
	var str string = "Go"      // Kiểu string
	var i int = 42             // Kiểu int
	var f float64 = 3.14       // Kiểu float64
	var c2 complex128 = 1 + 2i // Kiểu complex128
	fmt.Println(b2, str, i, f, c2)

	fmt.Println("-----------------------------------------------------")
	fmt.Println("Zero Values")
	var i4 int     // Giá trị zero của int là 0
	var f4 float64 // Giá trị zero của float64 là 0.0
	var b4 bool    // Giá trị zero của bool là false
	var s4 string  // Giá trị zero của string là ""

	fmt.Printf("%v %v %v %q\n", i4, f4, b4, s4)

	fmt.Println("-----------------------------------------------------")
	fmt.Println("Type Conversions")
	i5 := 42         // int
	f5 := float64(i) // Ép kiểu int sang float64
	u5 := uint(f)    // Ép kiểu float64 sang uint

	fmt.Println(i5) // In ra: 42
	fmt.Println(f5) // In ra: 42.0
	fmt.Println(u5) // In ra: 4

	fmt.Println("-----------------------------------------------------")
	fmt.Println("Type inference")
	i6 := 42           // int
	f6 := 3.142        // float64
	g6 := 0.867 + 0.5i // complex128

	fmt.Printf("i: %T\n", i6) // int
	fmt.Printf("f: %T\n", f6) // float64
	fmt.Printf("g: %T\n", g6) // complex128

	fmt.Println("-----------------------------------------------------")
	fmt.Println("Type inference")
	const Pi = 3.14
	const Truth = true
	const Greeting = "Hello, Go!"

	fmt.Println(Pi)       // In ra 3.14
	fmt.Println(Truth)    // In ra true
	fmt.Println(Greeting) // In ra "Hello, Go!"

	fmt.Println("-----------------------------------------------------")
	fmt.Println("LOOP For")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("-----------------------------------------------------")
	fmt.Println("LOOP For === while")
	i7 := 0
	for i7 < 5 { // Không có init và post statement
		fmt.Println(i7)
		i7++ // Thực hiện cập nhật trong khối lệnh
	}

	fmt.Println("-----------------------------------------------------")
	fmt.Println("IF")
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println("-----------------------------------------------------")
	fmt.Println("SWITCH")

	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	fmt.Println("-----------------------------------------------------")
	fmt.Println("DEFER")
	exampleDefer()

}

// Hàm cộng hai số nguyên
func add(x int, y int) int {
	return x + y
}

func addParamSampleType(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return "return: " + y, "return: " + x
}

func calc(x, y int) (sum, diff int) {
	sum = x + y
	diff = x - y
	return // Naked return, không cần đối số
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func exampleDefer() {
	defer fmt.Println("Đây là câu lệnh deferred đầu tiên")
	defer fmt.Println("Đây là câu lệnh deferred thứ hai")

	fmt.Println("Hàm example bắt đầu")
}
