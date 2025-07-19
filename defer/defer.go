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

	p := func(a int) {
		fmt.Println("ami", a)
	}

	defer p(result)

	fmt.Println("third", result)

	defer fmt.Println(5)

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
