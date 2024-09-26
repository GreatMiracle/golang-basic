package main

import "fmt"

// Định nghĩa struct ContactInfo
type ContactInfo struct {
	Email   string
	ZIPCode int
}

// Định nghĩa struct Person
type Person struct {
	FirstName string
	LastName  string
	Contact   ContactInfo // Nhúng struct ContactInfo vào struct Person
}

func main() {
	//Tạo một instance của Person với ContactInfo
	jim := Person{
		FirstName: "Jim",
		LastName:  "Party",
		Contact: ContactInfo{ // Tạo một instance của ContactInfo
			Email:   "jim@gmail.com",
			ZIPCode: 94000,
		},
	}

	//In ra thông tin của jim
	fmt.Printf("First Name: %s\nLast Name: %s\nEmail: %s\nZIP Code: %d\n",
		jim.FirstName, jim.LastName, jim.Contact.Email, jim.Contact.ZIPCode)
}
