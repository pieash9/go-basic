package main

import "fmt"

func main() {

	subtraction := func(a, b int) {
		fmt.Println(a - b)
	}
	subtraction(66, 45) // will work

	add(2, 3) // not work but in global it will work

	// function expression
	add := func(a, b int) {
		fmt.Println(a + b)
	}
}
