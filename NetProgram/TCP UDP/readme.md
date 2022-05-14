# Golang 的TCP和UDP实现

TCP和UDP存在于Application layer   

## TCP in Go
传输稳定， 需要用 net包,关于TCP 协议的支持涉及到 TCPAddr、TCPConn、TCPListener，大多数场景中，并不需要直接调用这些类型  
## UDP in Go
传输快，net包关于UDP协议类型包括UDPConn 和UDPAddr .多数例子是基于这个基础类型直接实现UDP。Go提供了很多抽象接口用于实现UDP网络通信， 最为重要的包是 **PacketConn 包**

# Server端常用函数、接口

#### Listen函数：
```
func Listen(network, address string) (Listener, error)
	network：选用的协议：TCP、UDP， 	如：“tcp”或 “udp”
	address：IP地址+端口号, 			如：“127.0.0.1:8000”或 “:8000”

```		
#### Listener 接口：
```
type Listener interface {
			Accept() (Conn, error)
			Close() error
			Addr() Addr
}

```  
#### Conn 接口：
```
type Conn interface {
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
	Close() error
	LocalAddr() Addr
	RemoteAddr() Addr
	SetDeadline(t time.Time) error
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
}

```