/**
TODO 接口都是以er 结尾


 */

package main

import "fmt"

type Reader interface {
	Read(str string)
}

type File struct {
	fileName string "filename"
	reader *Reader
}

func (file *File) Read(str string)  {
	fmt.Print("str...", str)
}

func run(reader *Reader)  {

}


/**
几个接口的小工作, client 有发送数据的能力，定义一个client 和一个发送接口 ，然后设置一个接口的实现，把client上设置上
TODO 这个证明了如果一个结构体实现了一个接口方法，这个结构体是可以设置到接口调用地方的，但如果方法大小写变了就不行了，http里面那个例子可以找到大写的方法实现的
两个接口可以实现的方法一摸一样，这个还是挺有意思的
接口里也不能包含变量
 */
type CustomTransporter interface {
	Send(str string)
}

type CustomTransporter2 interface {
	Send(str string)
}

type CustomClient struct {

	transport CustomTransporter
}

type HttpTransport struct {

}

func NewCustomClient(transport CustomTransporter) *CustomClient {
	cc := &CustomClient{transport:transport}
	return cc
}

func NewHttpTransport() *HttpTransport {
	ht := &HttpTransport{}
	return ht
}

func (ht *HttpTransport)Send(str string)  {
	fmt.Printf("http transport %v", str)
}

func (cc *CustomClient)call()()  {
	cc.transport.Send("why")
}

func main() {
	//readerQ := *new(Reader)
	//file := File{reader: &readerQ}
	//file.fileName = "xxx"
	//file.Read("yyy")
	//_ := file.reader
	//// 这行不可用
	////readerW.Read("xxx")

	ht := NewHttpTransport()

	cc := NewCustomClient(ht)

	cc.call()

}


