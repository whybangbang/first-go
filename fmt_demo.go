/**
https://golang.org/pkg/fmt/
这个包主要来学习fmt相关内容
 */
package main

import "fmt"

func fmt_test()  {

	fmt.Printf("%v\n", "aa")

	fmt.Printf("%#v\n", "aa")
	// 打印类型的
	fmt.Printf("%T \n", "aa")
	// TODO  这个不知道干嘛的
	fmt.Printf("%%\n")

	// 各种基本类型打印
	fmt.Printf()
	fmt.Fprint()

}

func main()  {
	fmt_test()
}
