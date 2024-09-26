package main

import "fmt"

func main() {
	// Khai báo map với các cặp key-value
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	// Truy cập giá trị qua key
	fmt.Println("Mã màu của red là:", colors["red"])

	// Thay đổi giá trị của key "blue"
	colors["blue"] = "#1E90FF"
	fmt.Println("Mã màu mới của blue là:", colors["blue"])

	fmt.Println(colors)

	var colorsNil map[string]string
	fmt.Println("colorsNil: ", colorsNil) // In ra map rỗng

	colorsMake := make(map[string]string)
	colorsMake["white"] = "#FFFFFF"
	colorsMake["blue1"] = "#ABCD"
	fmt.Println("colorsMake: ", colorsMake)

	delete(colorsMake, "blue1")
	fmt.Println("after action delete colorsMake: ", colorsMake)

	for colorKey, colorValue := range colors {
		fmt.Println("Hex code for", colorKey, "is", colorValue)
	}

}
