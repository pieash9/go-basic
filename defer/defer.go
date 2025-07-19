package deferExplain

import "fmt"

func calculate() (result int) {
	fmt.Println("First", result)

	show := func() {
		result = result + 10
		fmt.Println("defer", result)
	}
	defer show()

	result = 5
	fmt.Println("Second", result)

	return
}

func DeferExplain() {
	a := calculate()
	fmt.Println("calculate", a)
}

func calc() int {
	result := 0
	fmt.Println("First", result)

	show := func() {
		result = result + 10
		fmt.Println("defer", result)
	}
	defer show()

	result = 5
	fmt.Println("Second", result)

	return result
}

func Defer() {
	a := calculate()
	fmt.Println("calculate", a) // 15
	b := calc()
	fmt.Println("calc", b) // 5
}
