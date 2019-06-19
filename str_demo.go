/**
main 包必须package开头
 */
package main

import "fmt"

import (
	"strconv"
	"strings"
	"unicode"
)

const pi  = 3.1415
const country string  = "china"

func isLetter(c rune) bool {
	return !unicode.IsLetter(c)
}

func StringsTest(str string) {
	strList := strings.Split(str, ",")
	fmt.Println(strList)
	strCount := strings.Count("北京市北北", "北京")
	fmt.Println(strCount)

	// 因为这个返回的是数字，所以如果这个开启下面会报错
	//result := strings.Compare("a", "b")
	fmt.Println(strings.Compare("q", "1"))

	// 正常的包含关系，里面实现走的字节 IndexBytes
	result := strings.Contains("北京大小", "北京房")
	fmt.Println("Contain result:", result)

	// 是否右边的多个词里面有一个在左边，只要有一个则为true
	result = strings.ContainsAny("北京大小", "北京房")
	fmt.Println("ContainAny result:", result)
	// 这个是unicode code 91 代表 a
	result = strings.ContainsRune("a北京大小", 97)
	fmt.Println("ContainRune result:", result)

	// 理解为utf-8是否是相等的，这里面实现了ascii 码的大小写 a A
	fmt.Println(strings.EqualFold("北", "北"))

	// 中文空格和英文空格一样
	fmt.Println(len(strings.Fields("中国 xx 123, 呵呵 哈哈")))

	//传入一个功能来是否做区分
	fmt.Println(strings.FieldsFunc("中国 abc ,,,", isLetter))

	// 是否有前缀，这个直接slice判断的为什么不用byte
	result = strings.HasPrefix("中国人", "中国")
	fmt.Println(result)

	// 把一个slice 用sep 连接到一起
	sSlice := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(sSlice, "|"))

	// TODO 从这个方法可以看出 func也是可以作为参数类型的
	//strings.Map()

	//重复几次
	fmt.Println("ba", strings.Repeat("na", 2))

	// 替换，替换几个
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))

	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	// 和上面的区别就是这个是在逗号后面做分割
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))

	// 这个是首字母大写
	fmt.Println(strings.Title("!herroyal highnes"))
	// 这个是所有字母大写
	fmt.Println(strings.ToTitle("loud noises"))

	/**
	StringBuilder 系列，用这个可以stringbuilder一样减少数据复制，加快速度，但也一样需要自己去考虑指定 大小，
	TODO rune 因为unicode 所以可能会默认grow数据值，但writeString没有不知道为啥！！！
	 */

	strBuilder := strings.Builder{}
	strBuilder.WriteString("aaa")
	fmt.Println(strBuilder.String())

}

func StrConvDemo(str string)()  {
	// 字符串转int
	atoiVar, _ := strconv.Atoi("2")
	fmt.Println(atoiVar)

	// 浮点类型转换
	formatFloatVar := strconv.FormatFloat(1.23, 'f', 5, 32)
	fmt.Println(formatFloatVar)

	fmt.Println("int size:", strconv.IntSize)
}


var intA int8
var intB int16
var intC int32
var intD int64
var intE  uint8

func intOPera()  {
	intA = 23
	intB = int16(intA)

	intB = intB + int16(intA)
	intB = intB + 5
	fmt.Println(intA, intB)
}

var fA float32
var fB  float64


func main()  {
	//StringsTest("a,b")
	//StrConvDemo("ss")
	intOPera()
}





