// TCP实现并发-客户端

// 上面都是单机版的客户端通信，如果想要实现并发，需要使用Goroutine+循环实现

// 循环读取客户端发送的数据
// 如果客户端强制关闭连接需要做处理
// 客户端发送exit时

package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// 创建监听套接字
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen出错:", err)
		return
	}
	defer listen.Close()

	// 创建客户端连接请求
	fmt.Println("服务器启动成功，等待客户端连接！")
	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept出错:", err)
			return
		}
		// 调用服务器和客户端通信的函数
		go HandlerConnect(accept)
	}
}

func HandlerConnect(accept net.Conn) {
	defer accept.Close()

	// 获取客户端发送的数据
	// 获取连接客户端的网络地址
	addr := accept.RemoteAddr()
	fmt.Println(addr, "客户端连接成功！")

	buf := make([]byte, 4096)
	for {
		read, err := accept.Read(buf)
		if err != nil {
			fmt.Println("accept.Read出错:", err)
			return
		}
		fmt.Println("服务器读到数据:", string(buf[:read]))

		// 模拟服务器收到数据后，回发给客户端，小写转大写
		data := strings.ToUpper(string(buf[:read]))
		accept.Write([]byte(data))
	}
}
