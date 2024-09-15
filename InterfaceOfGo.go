package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Chuvi() float64
	Dientich() float64
}

type Hinhtron struct {
	Bankinh float64
}

type Hinhchunhat struct {
	Chieudai  float64
	Chieurong float64
}

func (t Hinhtron) Chuvi() float64 {
	return 2 * math.Pi * t.Bankinh
}

func (t Hinhtron) Dientich() float64 {
	return math.Pi * math.Pow(t.Bankinh, 2)
}

func (n Hinhchunhat) Chuvi() float64 {
	return (n.Chieurong + n.Chieudai) * 2
}

func (n Hinhchunhat) Dientich() float64 {
	return n.Chieudai * n.Chieurong
}

func main() {
	c := Hinhtron{5}
	r := Hinhchunhat{6, 4}

	//shapes := []Shape{c, r}

	var a Shape = c
	fmt.Println(a.Dientich())

	a = r
	fmt.Println(a.Dientich())

	// Tạo danh sách các hình (Shape)
	shapes := []Shape{c, r}
	for i, value := range shapes {
		fmt.Printf("Index: %d, gia_tri_khoi_tao: %+v , Chuvi: %.2f - Dientich: %v \n", i, value, value.Chuvi(), value.Dientich())
	}

}
