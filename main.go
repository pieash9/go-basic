package main

import "fmt"

var a = 100 // data segment

func call() { // code segment
	add := func(x, y int) {
		z := x + y
		fmt.Println(z)
	}
	add(5, 6)
	add(51, 61)

}

func main() { // code segment
	call()
	fmt.Println(a)
}

func init() { // code segment
	fmt.Println("Hello")
}
