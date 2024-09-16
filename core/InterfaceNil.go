package main

import "fmt"

// Định nghĩa interface
type Describer interface {
	Describe()
}

// Cấu trúc có tên Person
type Person struct {
	Name string
	Age  int
}

// Cài đặt phương thức Describe cho Person
func (p Person) Describe() {
	if p.Name != "" {
		fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
	} else {
		fmt.Printf("nilllllllllllll")
	}
}

type PersonPointer struct {
	Name string
	Age  int
}

// Cài đặt phương thức Describe cho Person
func (p *PersonPointer) Describe() {
	if p == nil {
		fmt.Println("Person is nil")
		return
	}
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

// Cấu trúc có tên Car
type Car struct {
	Brand string
	Year  int
}

// Cài đặt phương thức Describe cho Car
func (c Car) Describe() {
	fmt.Printf("Brand: %s, Year: %d\n", c.Brand, c.Year)
}
func main() {
	// Khai báo các đối tượng thuộc các kiểu khác nhau
	var d Describer

	// Đối tượng kiểu Person
	p := Person{Name: "John", Age: 30}
	d = p
	d.Describe() // Gọi phương thức Describe của Person

	// Đối tượng kiểu Car
	c := Car{Brand: "Toyota", Year: 2020}
	d = c
	d.Describe() // Gọi phương thức Describe của Car
	//
	//// Xử lý nil
	var pNil *PersonPointer
	d = pNil
	if d != nil {
		d.Describe() // Không gây lỗi, xử lý nil receiver
	} else {
		fmt.Println("Interface is nil")
	}

}
