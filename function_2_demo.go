package main

import "fmt"

// 递归
func fibonacci(a, b int)  {
	value :=  a + b
	if b > 100 {
		return
	}

	fmt.Println("fibonacci", a, b)
	fibonacci(b, value)
}

func callback(a, b int)  {
	fmt.Printf("invoke")
	fmt.Printf("cal%d, %d", a, b)
}

func cal(a ,b int, f func(int, int))  {
	f(a, b)
}

//回调

func main() {
	//fibonacci(1, 2)
	cal(1, 2, callback)
}