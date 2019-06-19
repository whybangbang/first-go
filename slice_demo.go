package main

import "fmt"

//声明数组,这样就初始化数组了，不过都是0
var firstArr [5]int
// TODO []string [3]string 这两类数据类型是不一样的，如果这个地方使用[3]string，则printArr不能使用了
var seArr  = []string{"one", "two", "three"}

func operation()  {

	firstArr[2] = 5

	for pos, value := range firstArr {
		fmt.Println(pos, value)
	}

	for i := 0; i < len(firstArr); i++ {
		fmt.Println(firstArr[i])
	}

}


func printArr(arr []string)  {
	for pos, value := range arr {
		fmt.Println(pos, value)
	}

}

func test()  {
	// 这个的数量由后面的初始值决定了
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
}

// 指针问题
func test2()  {
	// TODO arr1 的类型是 *[5]int 而 arr2 的类型是 [5]int
	var arr1 = new([5]int)
	arr1[1] = 10
	var arr2 [5]int
	// TODO 如果用arr2 = arr1 会报错 cannot use arr1 (type *[5]int) as type [5]int in assignment
	// arr3更改就更改了
	arr2 = *arr1

	arr3 := arr1

	arr2[2] = 100
	arr3[2] = 300

	fmt.Println(arr2)
	fmt.Println(arr1)
}

func test3(arr *[3]int)  {
	arr[2] = 3
}

var keyArr = [...]int{1:2, 3: 4}



/**
切片（slice）是对数组一个连续片段的引用
0 <= len(s) <= cap(s)
TODO &符代表获取该数据的引用
这里有图，slice 是在array上再做了一哥结构
https://learnku.com/docs/the-way-to-go/72-slice/3613
切片重组 是因为切片是数组的一部分，如果达到数组的最大容量需要扩容

 */

var firstSlice []int
var firstArray [5]int
var sendSlice = firstArray[1:3]
var thirdSlice = []int{2, 3, 5, 7, 11}[:]
var fourthSlice = make([]int, 50, 100)


func sliceDemo()  {
	// TODO 这个好像不行
	//ss := [3]int{1,2,3}[:]
	//fmt.Println(ss)
	fmt.Printf("%T \n", thirdSlice)

	for _, value := range thirdSlice {
		fmt.Println(value)
	}

	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}

func main() {
	//operation()

	//printArr(seArr)
	//test2()

	//var aaa [3]int
	//test3(&aaa)

	//var arr4 = new([3]int)
	//test3(arr4)
	//fmt.Println(arr4)
	sliceDemo()


}