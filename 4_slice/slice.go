package slice

import "fmt"

func print(numbers ...int) { // slice hisebe receiver (variadic parameter)
	fmt.Println(numbers)      // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println(len(numbers)) // 10
	fmt.Println(cap(numbers)) // 10
}

func changSlice(p []int) []int {
	p[0] = 10         // [10, 6, 7] ekhne purber 0 index er value 5 ke change kore 10 krbe
	p = append(p, 11) // [10, 6, 7, 11] len= 4, cap = 6
	return p
}

func Slice() {
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6) // ptr = 25, len = 6, cap = 10
	x = append(x, 7) // ptr = 25, len = 7, cap = 10

	a := x[4:] // 4 index theke sob gula value niye toiri krbe [5, 6, 7] ptr = 29, len = 3, cap = 6 (slice kora value gulo cap theke bad jabe)

	y := changSlice(a) // [10, 6, 7, 11] ptr = 29, len = 4, cap = 6

	fmt.Println(x)      // [1 2 3 4 10 6 7]
	fmt.Println(y)      // [10 6 7 11]
	fmt.Println(x[0:8]) // [1 2 3 4 10 6 7 11]

	print(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

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
