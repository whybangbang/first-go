package main

import "fmt"

// TODO 申明一个外部定义的函数
//func flushCache(begin uintptr) int

//命名返回参数，只需要在函数中使用到该名字即可
// int32 不能 * ？ TODO
func name1(a ,b int) (x, y int)  {
	x = a * b
	y = a + b
	return 
}
/**
指针传递
传递指针要比传递副本要消耗少
slice map interface 默认指针传递，函数默认副本传递
TODO 这个地方有疑点
*/
func multiValue(a *int)  {
	*a = 5
}

/**
可变参数
如果类型也不确定可以采用 struct 或者interface来处理
 */
func min(a ...int) (int) {

	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}

	sliceTransfer(a)

	return min
}

// 可变参数可以作为slice传递
func sliceTransfer(a []int)  {
	
}

// defer 类似finally 做收尾工作，但这玩意儿为啥要在中间可以设置？

func main() {
	//name1(1, 2)
	//
	//// 这就是指针操作
	//var p int
	//multiValue(&p)
	//fmt.Println("p:",p)

	value := min(5,4,5,6,7)
	fmt.Println(value)
}
