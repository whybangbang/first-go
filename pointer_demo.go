package main

import "fmt"

/**
创建指针用 *  获取一个数的指针用&
& 是取地址符
* 指针获取真实值用这个

声明一个指针，如果没有设置值需要用new分配内存,分配内存之后就设置值了
 */

var pointerOne *int


/**
new 是返回一个值的指针，而且设置为了默认值
make 只能用于slice map channel，返回这个值的引用
 */

func main() {

	fmt.Println(pointerOne)
	pointerOne = new(int)
	fmt.Println(pointerOne, *pointerOne)

	aa := 5

	fmt.Println(pointerOne)
	pointerOne = &aa

	fmt.Println(pointerOne, &aa)
	fmt.Println(*pointerOne)
	*pointerOne = 6
	fmt.Println(*pointerOne)
}