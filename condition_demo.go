package main

import "fmt"

func ifELse(condition, condition2 bool)  {
	if true{
		fmt.Println("is true");
	}

	if condition {
		fmt.Printf("%v\n", "condition is true")
	} else if condition2 {
		fmt.Printf("%v\n", "condition 2 is true")
	} else {
		fmt.Printf("%v\n", "c1 c2 all false")
	}

}

func switchDemo()  {

	// TODO fallthrough 是直接就会执行后面的代码，儿没有判断!!!
	k := 6
	switch k {
		case 4:
			fmt.Printf("%v\n", 4);
			fallthrough;
		case 5:
			fmt.Printf("%v\n", 5);
			fallthrough;

		case 6:
			fmt.Printf("%v\n", 6);
			fallthrough;

		case 7:
			fmt.Printf("%v\n", 7);
			fallthrough;

		default:
			fmt.Printf("%v\n", "default")



	}
}

func forDemo()  {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v ", i);
	}

	var a int8
	a = 5
	for a > 0 {
		fmt.Printf("while loop %v\n", a)
		a -= 1
	}

	// range
	aStr := "abcdefg"
	for pos, i := range aStr {
		fmt.Printf("%v %d \n", pos, i)
	}

	// goto 来模拟循环
	gI := 0
	LABEL1:
		fmt.Printf("%v %d \n", gI, gI)
		gI++

		if gI == 5{
			return;
		}
		goto LABEL1
}

func main()  {
	//ifELse(false, false)
	forDemo()
}
