package main

import (
	"bufio"
	"fmt"
	"net"
)

// TCP服务端程序的处理流程：

//     1.监听端口
//     2.接收客户端请求建立链接
//     3.创建goroutine处理链接。

// TCP server port

// 处理函数
func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client  failed, err: ", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("Data from client:", recvStr)
		recvDetails := "input is " + recvStr
		_, err = conn.Write([]byte(recvDetails)) // 发送数据
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err: ", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err: ", err)
			continue
		}
		go process(conn) // 启动一个gorountine
		fmt.Println("server is connecting client.")
	}
}
