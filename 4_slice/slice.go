package slice

import "fmt"

func Slice() {
	arr := [6]string{"this", "is", "a", "go", "interview", "Ques"}
	fmt.Println(arr)
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
