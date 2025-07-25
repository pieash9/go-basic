package goroutine

import (
	"fmt"
	"time"
)

var a = 10
var p = 11

func printHello(num int) {
	time.Sleep(5 * time.Second) // 5 second sleep to simulate work
	fmt.Println("Hello", num)
}

func GO_Routine() {
	var x int = 10
	fmt.Println(x)
	go printHello(1) // 5 ti alada go routine e run hbe
	go printHello(2) // 5 ti alada go routine e run hbe
	go printHello(3) // 5 ti alada go routine e run hbe
	go printHello(4) // 5 ti alada go routine e run hbe
	go printHello(5) // 5 ti alada go routine e run hbe
	fmt.Println(a, " ", p)

	time.Sleep(5 * time.Second)
	fmt.Println("End of GO_Routine")
}
