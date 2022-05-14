// UDP实现并发 Server

// // Create Listen Address 创建监听地址
// func ResolveUDPAddr(network, address string) (*UDPAddr, error)

// // Create Listen UDP  创建监听连接
// func ListenUDP(network, address string) (*UDPConn, error)

// // 接收UDP数据
// func (c *UDPConn) ReadFromUDP(b []btye) (int, *UDPAddr, error)

// // 写出数据到UDP
// func (c *UDPConn) WriteToUDP(b []btye) (int, *UDPAddr, error)

// e.g.
package main

import (
	"fmt"
	"net"
)

func main() {
	// 指定服务器的ip和端口，和TCP协议不一样，需要先写好再传给ListenUDP使用
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.ResolveUDPAddr err：", err)
		return
	}

	fmt.Println("服务器启动成功！")
	// 创建用户通信的Socket
	udpConnect, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		fmt.Println("net.ListenUDP err：", err)
		return
	}
	defer udpConnect.Close()

	fmt.Println("服务器创建Socket成功！")

	// 读写客户端的数据
	buf := make([]byte, 4096)
	count := 0
	for {
		// 返回值：n int（读到的字节数）， addr *UDPAddr（客户端的地址）， err error
		udpBytes, ConnectAddr, err := udpConnect.ReadFromUDP(buf)

		if err != nil {
			fmt.Println("udpConnect.ReadFromUDP err：", err)
			return
		}
		count++
		// 模拟处理数据
		fmt.Printf("服务器读到第%v条数据 %v ：%s\n", count, ConnectAddr, string(buf[:udpBytes]))

		go func() {
			// 回写数据到客户端
			udpConnect.WriteToUDP([]byte("回写数据到客户端\n"), ConnectAddr)
		}()
	}
}
