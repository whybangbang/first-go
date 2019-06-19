package socket_demo

import (
	"fmt"
	"io/ioutil"
	"net"
)

/**
https://astaxie.gitbooks.io/build-web-application-with-golang/en/08.1.html

客户端 用DialTCP 建立 TCPConn 然后开始通信，直到客户端和服务端有一端关闭
 */



func main() {

	var remoteAddress *net.TCPAddr
	var rError error

	var tcpConn *net.TCPConn

	remoteAddress, rError = net.ResolveTCPAddr("tcp4", "127.0.0.1:9999")

	if rError != nil {
		fmt.Print("error")
	}
	tcpConn, rError = net.DialTCP("tcp", nil, remoteAddress)

	//tcpConn.Read()
	tcpConn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))

	ioutil.ReadAll(tcpConn)
}
