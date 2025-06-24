package main

import "fmt"

func call() func(x int, y int) {
	return add
}

func add(a int, b int) {
	c := a + b
	fmt.Println(c)
}

func main() {
	sum := call() // function expression
	sum(7, 3)
}
