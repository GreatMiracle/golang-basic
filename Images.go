package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// CustomImage là một loại ảnh tùy chỉnh
type CustomImage struct{}

// ColorModel trả về mô hình màu của ảnh
func (img CustomImage) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds trả về kích thước của ảnh
func (img CustomImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100) // Ảnh có kích thước 100x100
}

// At trả về màu đỏ tại mọi tọa độ
func (img CustomImage) At(x, y int) color.Color {
	return color.RGBA{255, 0, 0, 255} // Màu đỏ
}

func main() {
	m := CustomImage{}

	// Tạo một file để lưu ảnh PNG
	f, _ := os.Create("red_image.png")
	defer f.Close()

	// Lưu ảnh dưới dạng PNG
	png.Encode(f, m)
}
