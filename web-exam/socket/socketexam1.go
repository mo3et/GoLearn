// Articles:https://juejin.cn/post/6844903575932370952

package main

import "net"

// TODO Socket 是对TCP、UDP封装后提供的一层接口。可以利用Socket编写server和client(服务端和客户端)。然后让client和server建立TCP或UDP的连接

// Unix Socket 编程的函数接口
// 主要调用 listen、accept、write、read 等函数实现的

// Golang 中的Socket编程模型
// 和Linux Socket 编程相比，Go的Socket编程是
// Server(服务端)直接通过Listen + Accept 模式即可实现

// Socket Program model in Go

// Server  Use Listen + Accept Mode
func connHandler(c net.conn) {
	server, err := c.Read(buf)
	c.Write(buf)
}
func main() {
	server, err := net.Listen("tcp", ":1208")
	for {
		conn, err := server.Accept()
		go connHandler(conn)
	}
}

// Client Dial

func connHandler(c net.Conn) {
	for{
		c.Write(...)
		c.Read(...)
	}
}

func main() {
	conn,err:=net.Dial("tcp","localhost:1208")
	connHandler(conn)
}


// 实现一个可以接受不同命令的server

// 在同Folder的另外蓝光 server 与client