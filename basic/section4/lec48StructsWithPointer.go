package main

import "fmt"

// Định nghĩa struct Person
type Person struct {
	firstName string
	lastName  string
}

// Hàm để cập nhật tên, sử dụng con trỏ để cập nhật giá trị thực
func (p *Person) updateName(newFirstName string) {
	//(*p).firstName = newFirstName
	p.firstName = newFirstName
}

// Hàm để in thông tin của đối tượng
func (p Person) print() {
	fmt.Printf("First Name: %s, Last Name: %s\n", p.firstName, p.lastName)
}

func main() {
	// Khởi tạo đối tượng Jim
	jim := Person{
		firstName: "Jim",
		lastName:  "Beam",
	}

	// In thông tin ban đầu
	jim.print()

	// Sử dụng hàm updateName để cập nhật tên (thông qua con trỏ)
	jim.updateName("Quan.Le")

	// In thông tin sau khi cập nhật
	jim.print()
}
