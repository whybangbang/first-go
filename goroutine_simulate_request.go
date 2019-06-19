/**
模拟多路复用的请求

这个设计挺有意思,其实就是都交给chan处理
请求的时候有哪些地方需要处理
(1) 首先请求需要都发在chan中，所以chan中是request对象
(2) 然后处理完之后又都丢回chan中，这个地方我感觉不用chan也行

关键是接受数据一个 go 然后处理一个go

 */
package main

import (
	"fmt"
	"time"
)

type Reply struct {
	respId int
}

func NewReply(respId int) *Reply {
	return &Reply{respId:respId}
}

type Request struct {
	reqId int
	replyc chan *Reply
}

// TODO 结构体里面的chan得初始化，否则会停止不了？为啥
func NewRequest(reqId int) *Request {
	return &Request{reqId:reqId, replyc:make(chan *Reply)}
}

func exec(reqC chan *Request)  {
	req := <- reqC
	//fmt.Println("send to response")
	req.replyc <- NewReply(req.reqId)
}

func startServer() chan *Request{
	requestC := make(chan *Request)
	go func() {
		for{
			// TODO 这个for 在主线程关闭的时候其实是没有关闭的，可以单独创建一个信号量的channel来关闭
			go exec(requestC)
		}
	}()

	return requestC
}

func main() {
	// start 一个server 接受
	requestC := startServer()

	var reqs [100] *Request
	for i := 0; i < 100; i++ {
		request := NewRequest(i)
		reqs[i] = request
		requestC <- request
	}

	for _, request := range reqs{
		replyc := request.replyc
		fmt.Println(*<-replyc)
	}




	time.Sleep(1e9)
}


