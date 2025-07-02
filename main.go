package main

import "fmt"

type User struct { // code segment e thakbe => type User custom type
	Name string // member variable or property
	Age  int    // member variable or property
}

func main() {
	var user1 User

	user1 = User{ // instance or Object => instance bananor prokriyake bole instantiate
		Name: "Pieash",
		Age:  25,
	}

	user2 := User{ // instance or Object
		Name: "Sheikh",
		Age:  28,
	}

	fmt.Println("Name:", user1.Name)
	fmt.Println("Age:", user1.Age)
	fmt.Println("Name:", user2.Name)
	fmt.Println("Age:", user2.Age)
}
