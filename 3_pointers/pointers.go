package pointers

import "fmt"

func print(numbers [3]int) {
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
	print(arr)
}

/*
	2. Phases:
		1. compilation phase (compile time)
		2. execution phase (run time)

  *************** compile phase ************
  ** code segment **
	main = func () {...}

*/
