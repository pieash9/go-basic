package receiver_func

import "fmt"

type User struct { // code segment e thakbe => type User custom type
	Name string // member variable or property
	Age  int    // member variable or property
}

func printUserDetails(usr User) {
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}

func (usr User) printDetails() { // receiver function
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}

func (usr User) call(a int) { // receiver function
	fmt.Println(usr.Name)
	fmt.Println(a)
}

func Receiver() {
	var user1 User

	user1 = User{ // instance or Object => instance bananor prokriyake bole instantiate
		Name: "Rahim",
		Age:  25,
	}

	user2 := User{ // instance or Object
		Name: "Sheikh",
		Age:  28,
	}

	user1.printDetails() // receiver function call
	user1.call(10)
	printUserDetails(user2)
}
