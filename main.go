package main

import "fmt"

func add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	add(5, 6)
	add(51, 61)
}

func init() {
	fmt.Println("Hello")
}
