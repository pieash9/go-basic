package slice

import "fmt"

func Slice() {
	var x []int      // [], ptr = nil, len = 0, cap = 0
	x = append(x, 1) // [1], ptr = 26 ,len = 1, cap = 1 (heap e rakhbe value)
	x = append(x, 2) // [1, 2], ptr = 27 ,len = 2, cap = 2 (heap e rakhbe value)  // capacity & length same hole 2 dara gun krbe
	x = append(x, 3) // [1, 2, 3], ptr = 29 ,len = 3, cap = 3 (heap e rakhbe value)  // capacity & length same hole 2 dara gun krbe
	y := x

	x = append(x, 4)
	y = append(y, 5)

	x[0] = 10

	fmt.Println(x) // [10 2 3 5]
	fmt.Println(y) // [10 2 3 5]
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
	so here arr[1:4] => ptr=1, len=3(je koita value ache), cap=5(jekhane thke shuru hoise ekbre shesh prjnto)

	}

	7. slice underlying array rule => 1024 * 2 = 2048 -> double kore parbe capacity.. but 1024 par hoye gele 25% kore barbe.
*/
