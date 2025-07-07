package pointers

import "fmt"

// pass by value => direct value pathano ekhtre copy hbe ekti ekti kore
// pass by reference => starting point er address pass of array

type User struct {
	Name   string
	Age    int
	Salary float64
}

func print(numbers *[3]int) {
	fmt.Println(numbers)
}

func Pointers() {
	// pointer or address of memory (ram or hard disk)
	// x := 20
	// p := &x // ampersand & => address of 0xc00000a0d8 (hexa format, 0x means hexa)| 82,463,37,62,008(decimal format)

	// fmt.Println("Value:", *p)
	// *p = 30                                  // change the value by the address
	// fmt.Println("Address:", p)               // p is the address of x
	// fmt.Println("value at the address:", *p) // *p is the value at the address | * is value at address

	arr := [3]int{1, 2, 3}
	print(&arr)

	obj := User{ // instance or obj
		Name:   "Pieash",
		Age:    26,
		Salary: 0,
	}

	p := &obj // address of obj
	fmt.Println(*p)

	fmt.Println(obj.Age)
}

/*
	2. Phases:
		1. compilation phase (compile time)
		2. execution phase (run time)

  *************** compile phase ************
  ** code segment **
  print = func(numbers int[3]) {...}
	main = func () {...}

*/
