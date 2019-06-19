package main

import "fmt"

type struct1 struct {
	i1  int "tag"
	f1  float32
	str string
}

// 这是第一种创建结构体的方法，分配内存，拿到的是指针
func first()  {
	ms := new(struct1)
	fmt.Println(*ms)
}

type struct2 struct {
	i int
}

var a  struct2
var b  *struct2

func sencond()  {
	fmt.Println(a.i)
	//b = new(struct2)
	//b.i = 2
	fmt.Println(b.i)
}

func third()  {
	// 这是指针
	ms := &struct2{i: 4}
	fmt.Println(*ms)

	// 这是直接变量
	mx := struct2{i: 5}
	fmt.Println(mx)
}

// TODO 结构体中可以通过值来调用，也可以通过引用调用都是合法的

type User struct {
	age int
	name string
}

// 这种复制了一个对象
func upUser(user User)  {
	user.age = 5
	user.name = "up"
}

func upUserPointer(user *User)  {
	user.age = 5
	user.name = "up"
}

func upUserAction()  {
	//user := User{}
	user := &User{}
	fmt.Println("upUserAction before", user)
	upUser(*user)
	fmt.Println("upUserAction after", user)

	upUserPointer(user)
	fmt.Println("upUserAction after pointer", user)
}


// TODO 构造器工厂,必须以new 或者New开头
//
func NewUser(age int, name string) *User {
	//return &User{age: age, name: name}
	user := new(User)
	user.name = name
	user.age = age
	return user
}

// 结构体可以匿名

//结构体的方法 就是函数前面加上 t *T

/**
TODO 变量创建有三个过程

声明一个变量 或者声明一个指针 var aa type
声明指针这里这个的值是nil，如果没有分配内存操作是不鞥呢操作的！！！
分配一块内存 并设置初始化值 aa = new(type) or aa = &type



slice map  channel得单独考虑
 */

func main() {

	////first()
	////
	sencond()
	////
	////third()
	//
	////upUserAction()
	//
	//user := NewUser(66, "old")
	//fmt.Println(user)
	//upUserPointer(user)
	//fmt.Println(user)


}