package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	dial, err := net.Dial("udp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial出错：", err)
		return
	}
	defer dial.Close()

	for {
		// 发送数据
		dial.Write([]byte("我是客户端"))

		// 接收服务器返回的数据
		buf := make([]byte, 4096)
		read, err := dial.Read(buf)
		if err != nil {
			fmt.Println("accept.Read出错:", err)
			return
		}

		// 接收数据后处理数据
		fmt.Println("客户端获取到：", string(buf[:read]))
		time.Sleep(time.Second)
	}
}
