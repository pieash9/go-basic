package slice

import "fmt"

func Slice() {
	arr := [6]string{"this", "is", "a", "go", "interview", "Ques"}
	fmt.Println(arr)

	s := arr[1:4]
	fmt.Println(s)
}

/*
	2. Phases:
		1. compilation phase (compile time)
		2. execution phase (run time)

  *************** compile phase ************
  ** code segment **
  print = func(numbers int[3]) {...}
	main = func () {...}

	{
**	slice => Ptr=1, len=3, cap=5
** Array => arr := [6]string{"this", "this", "is", "a", "go", "interview", "Ques"}
	so here arr[1:4] => Ptr=1, len=3, cap=5

	}

*/
