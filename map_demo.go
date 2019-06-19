package main

import "fmt"

// key 可以是任意可以用 == 或者！= 操作符比较的类型，比如 string、int、float。所以数组、切片和结构体不能作为 key 
var map1 map[string]int
var mapLit = map[string]int{"one": 1, "two": 2}
// TODO 永远不要用new来声明map！！！
var mapCreated = make(map[string]float32)
// 如果知道容量最好指定， 对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明。
var mapCreated2 = make(map[string]float32, 100)

func operationMap()  {
	fmt.Println(mapLit["one"])
	mapLit["three"] = 3
	fmt.Println(mapLit)

	mp1 := make(map[int][]int)
	mp2 := make(map[int]*[]int)

	// 得创建slice 不能创建数组
	aaSlice := []int{1,2,3}[:]
	mp1[1] = aaSlice
	mp2[1] = &aaSlice

	fmt.Println(mp1)
	fmt.Println(mp2)
	fmt.Println(mp2[1])
}

func main() {
	operationMap()
}

