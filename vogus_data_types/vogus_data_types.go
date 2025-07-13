package vogusdatatypes

import "fmt"

func Vogusdatatypes() {
	var a int16 = 500
	var b int8 = 127  // int8 is the set of all signed 8-bit integers. Range: -128 through 127.
	var c uint8 = 255 // uint8 is the set of all unsigned 8-bit integers. Range: 0 through 255
	var d uint64 = 66
	var e, f float32 = 3.14, 1.59     // 32 bit ke half kore use krbe
	var g float64 = 2.718281828459045 // 32 + 32 2 ta merge kre use krbe
	var flag bool = true              // 8 bits => 1 byte
	r := '❤'
	var s string = "Myself pieash"
	fmt.Println(r)
	fmt.Printf("%c\n", r)    // %c is used for printing a single character
	fmt.Printf("%d\n", a)    // %d is used for decimal integers
	fmt.Printf("%.2f\n", e)  // %f is used for floating point numbers with 2 decimal places -> .2 means 2 decimal places
	fmt.Printf("%v\n", flag) // %v is used for printing the value of a variable in its default format
	fmt.Printf("%s\n", s)    // %s is used for printing strings
	fmt.Printf("%T\n", s)    // %T is used for printing the type of a variable
	fmt.Println(a, b, c, d, e, f, g, flag)

	// byte => alias for uint8 -> 8 bits per character -> 1 byte
	// rune => alias for int32 (unicode point) -> 32 bits per character -> 4 bytes -> %c
}
