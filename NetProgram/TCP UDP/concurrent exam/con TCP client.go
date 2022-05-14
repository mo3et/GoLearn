// TCP 实现并发

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//主动发送连接请求
	dial, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("et.Dial出错了", err)
		return
	}
	defer dial.Close()

	// os.Stdin()：获取用户键盘录入，
	go func() {
		str := make([]byte, 4096)
		for {
			read, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("os.Stdin.Read出错了", err)
				continue
			}

			// 读到的数据写给服务器，读多少写多少
			dial.Write(str[:read])
		}
	}()

	buf := make([]byte, 4096)
	// 回显服务器发送的数据，转成大写
	for {
		read, err := dial.Read(buf)

		// read=0的说明对端关闭连接，如果关闭连接这里就不需要往下读数据了
		if read == 0 {
			fmt.Println("检测到服务端端已经断开连接！")
			return
		}

		if err != nil {
			fmt.Println("回显服务器发送的数据dial.Read出错了", err)
			return
		}
		fmt.Println("客户端读到服务器的回显数据", string(buf[:read]))
	}
}
