package main

import "fmt"

func add(a, b int) {
	fmt.Println(a + b)
}

func main() {
	add(1, 2) // invoked the add func here

	a := 10    // expression
	if a > 0 { // if expression
		fmt.Println("a is greater than 0")
	}

	// Anonymous function
	func(a, b int) { // IIFE - Immediately Invoked Function Expression
		fmt.Println(a+b, "IIFE")
	}(7, 9)
}

func init() {
	fmt.Println("init")
}
