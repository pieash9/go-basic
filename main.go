package main

import "fmt"

var a = 100

func call() {
	add := func(x, y int) {
		z := x + y
		fmt.Println(z)

	}
	add(5, 6)
	add(51, 61)

}
func main() {
	call()
	fmt.Println(a)
}

func init() {
	fmt.Println("Hello")
}
