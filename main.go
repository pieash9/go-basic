package main

import "fmt"

const a = 10 // constant
var p = 100

func outer() func() {
	money := 100
	age := 30
	fmt.Println("Age = ", age)

	show := func() {
		money = money + a + p
		fmt.Println(money)
	}

	return show
}

func call() {
	incr1 := outer() // show func
	incr1()
	incr1() // second call e closure theke money er value nibe. & a , p er value global theke nibe

	incr2 := outer()
	incr2() // 210
	incr2() // 320
	incr2() // 430
}

func main() {
	call()
}

func init() {
	fmt.Println("====Bank====")
}
