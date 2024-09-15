package main

import "fmt"

// Node đại diện cho một nút trong danh sách liên kết
type Node[T any] struct {
	value T        // Giá trị của nút
	next  *Node[T] // Con trỏ tới nút tiếp theo
}

// List đại diện cho danh sách liên kết đơn
type List[T any] struct {
	head *Node[T] // Con trỏ tới nút dc gọi
}

// Add thêm một nút vào đầu danh sách
func (l *List[T]) Add(value T) {
	newNode := &Node[T]{value: value}
	newNode.next = l.head
	l.head = newNode
}

// PrintInOrder in các giá trị của danh sách
func (l *List[T]) PrintInOrder() {
	current := l.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func main() {

	intList := List[int]{}
	intList.Add(10)
	intList.Add(20)
	intList.Add(30)
	fmt.Println("Danh sách số nguyên:")
	intList.PrintInOrder()

}
