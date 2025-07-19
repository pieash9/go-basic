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

	result = 10
	p := func(a int) {
		defer fmt.Println(result)
		defer fmt.Println(a)
		fmt.Println("second", a)
		defer fmt.Println("third", a)
	}
	p(20)

	return
}

func DeferExplain() {
	a := calculate()
	fmt.Println("calculate", a)
}
