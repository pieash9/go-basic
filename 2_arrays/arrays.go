package arrays

import "fmt"

var (
	arr2 = [3]string{"I", "am", "Rahim"}
)

func Arrays() {
	var arr [2]int
	arr[0] = 5
	arr[1] = 10
	fmt.Println(arr)

	arr1 := [2]int{3, 6}
	fmt.Println(arr1)
	fmt.Println(arr2)

	fmt.Println(arr2[2])
}

/*
	2. Phases:
		1. compilation phase (compile time)
		2. execution phase (run time)

  *************** compile phase ************
  ** code segment **
	main = func () {...}

*/
