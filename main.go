package main

import "fmt"

func processOperation(a int, b int, op func(p int, q int)) {
	op(a, b)
}

func add(a int, b int) {
	c := a + b
	fmt.Println(c)
}

func main() {
	processOperation(6, 8, add)
}
