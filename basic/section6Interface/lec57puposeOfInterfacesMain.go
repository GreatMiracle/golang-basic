package main

import (
	"cards/section6Interface/lec57"
	"fmt"
)

// Hàm in lời chào
func PrintGreeting(g lec57.Greeter) {
	fmt.Println(g.GetGreeting())
}

func main() {
	eb := lec57.EnglishBot{}
	sb := lec57.SpanishBot{}

	PrintGreeting(eb) // In: Hello there.
	PrintGreeting(sb) // In: Hola.
}
