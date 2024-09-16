package main

import "fmt"

func main() {

	fmt.Println("-----------------------------------------------------")
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>Slice Literals, Slice Defaults <<<<<<<<<<<<<<<<<<<<<<<<<<<")
	// Slice literal với các giá trị boolean
	s := []bool{true, true, false}
	fmt.Println("Slice:", s) // In ra: [true true false]

	// Mảng với 10 phần tử kiểu int
	//var a [10]int
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Cắt slice với các giới hạn mặc định
	slice1 := a[:3] // Lấy 5 phần tử đầu tiên
	slice2 := a[5:] // Lấy phần tử từ vị trí 5 đến hết
	slice3 := a[:]  // Lấy toàn bộ mảng

	fmt.Println("Slice1:", slice1) // [0 0 0 0 0]
	fmt.Println("Slice2:", slice2) // [0 0 0 0 0]
	fmt.Println("Slice3:", slice3) // [0 0 0 0 0 0 0 0 0 0]

	// Độ dài và dung lượng của slice
	fmt.Println("Length of slice1:", len(slice1))   // Độ dài: 5
	fmt.Println("Capacity of slice1:", cap(slice1)) // Dung lượng: 10

	// Mở rộng slice1 nếu dung lượng cho phép
	slice1 = a[:7]                                 // Slice giờ chứa 7 phần tử
	fmt.Println("Slice1 sau khi mở rộng:", slice1) // [0 0 0 0 0 0 0]

	// Nếu mở rộng vượt quá dung lượng sẽ gây lỗi
	// slice1 = a[:11]  // Lỗi: ngoài phạm vi mảng

	fmt.Println("-----------------------------------------------------")
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>ARRAYS<<<<<<<<<<<<<<<<<<<<<<<<<<<")
}
