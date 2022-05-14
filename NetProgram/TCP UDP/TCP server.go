//TCP 服务端实现

package main

import (
	"fmt"
	"net"
)

func main() {
	// 指定服务器的通讯协议、ip、端口，Listen本身不做监听，这一步是创建了一个用于监听的Socket
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen出错：", err)
		return
	}
	defer listen.Close()

	fmt.Println("服务器启动完毕，等待客户端连接")

	// 阻塞监听客户端连接请求,成功建立连接后会返回用于通信的Socket
	accept, err := listen.Accept()
	if err != nil {
		fmt.Println("listen.Accept出错：", err)
		return
	}
	defer accept.Close()

	fmt.Println("服务器与客户端连接成功")

	// 读取客户端发动的请求
	buf := make([]byte, 4096)
	read, err := accept.Read(buf)
	if err != nil {
		fmt.Println("accept.Read出错：", err)
		return
	}

	// 接收数据后处理数据
	fmt.Println("服务器获取到：", string(buf[:read]))
}
