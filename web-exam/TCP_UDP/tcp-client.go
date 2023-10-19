// TCPAddr、TCPConn、TCPListener

package main

import (
	"fmt"
	"net"
)

//Connect TCP
// conn, err := net.Dial("tcp", "host:port")
// if err != nil {
// 	return err
// }
// defer conn.Close()

//simple Read
// buffer := make([]byte, 1024)
// conn.Read(buffer)

//simple write

// conn.Write([]byte("Hello from client"))

// TCP Client
func main() {
	// 指定用户端的通讯协议、ip、端口，Listen本身不做监听，这一步是创建了一个用于监听的Socket
	dial, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial出错：", err)
		return
	}
	defer dial.Close()

	// 发送数据
	dial.Write([]byte("我是客户端"))

	// 接收服务器返回的数据
	buf := make([]byte, 4096)
	read, err := dial.Read(buf)
	if err != nil {
		fmt.Println("accept.Read出错：", err)
		return
	}

	// 接收数据后处理数据
	fmt.Println("客户端获取到：", string(buf[:read]))
}
