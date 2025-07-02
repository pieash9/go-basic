package main

import "fmt"

type User struct { // code segment e thakbe => type User custom type
	Name string // member variable or property
	Age  int    // member variable or property
}

func printUserDetails(usr User) {
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}

func main() {
	var user1 User

	user1 = User{ // instance or Object => instance bananor prokriyake bole instantiate
		Name: "Pieash",
		Age:  25,
	}
	printUserDetails(user1)

	user2 := User{ // instance or Object
		Name: "Sheikh",
		Age:  28,
	}
	printUserDetails(user2)
}
