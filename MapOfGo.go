package main

import "fmt"

func main() {
	// Tạo một map dùng map literals
	ages := map[string]int{
		"Alice": 23,
		"Bob":   30,
		"Eve":   35,
	}

	// Truy cập phần tử từ map
	fmt.Println("Age of Bob:", ages["Bob"])

	// Thêm hoặc cập nhật phần tử
	ages["Alice"] = 25
	fmt.Println("Updated age of Alice:", ages["Alice"])

	// Xóa một phần tử
	delete(ages, "Eve")
	fmt.Println("Map after deleting Eve:", ages)

	// Kiểm tra sự tồn tại của một key
	age, ok := ages["Eve1"]
	if ok {
		fmt.Println("Age of Eve:", age)
	} else {
		fmt.Println("Eve not found in map")
	}
}
