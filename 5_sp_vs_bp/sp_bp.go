package spvsbp

import "fmt"

func add(x, y int) int {
	var result int
	result = x + y
	return result
}

func SP_vs_BP() {
	var a int = 10
	var sum = add(a, 4)
	fmt.Println(sum)
}
