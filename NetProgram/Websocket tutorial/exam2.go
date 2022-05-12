// socket或websocket连接的封装

// 发现只要是socket连接，都进行了封装，一种说法是为了线程安全，一种说法是为了业务的拓展性（比如说在发之前过滤一些消息）

package core

import "github.com/gorilla/websocket"

type Connection struct {
	Conn     *websocket.Conn
	inchan   chan []byte
	outchan  chan []byte
	exitchan chan bool
	isClose  bool
}

func NewConnection(conn *websocket.Conn) (*Connection, error) {
	wsConn := &Connection{
		Conn:     conn,
		inchan:   make(chan []byte, 1000),
		outchan:  make(chan []byte, 1000),
		exitchan: make(chan bool, 1),
	}
	return wsConn, nil
}

func (conn *Connection) Start() {
	// 不停的读长连接的消息，放到inchan中
	go conn.readLoop()
	// 不停的从outchan,读取消息，发到客户端
	go conn.writeLoop()
}

func (conn *Connection) ReadMessage() ([]byte, err error){
	if isClose {
		return nil,errors.New("connection is closed")
	}
	data = <-conn.inchan
	return data,nil
}

func (conn *Connection) WriteMessage(data []byte) error{
	if isClose{
		return nil,errors.New("connection is closed")
	}
	data=<-conn.inchan
}

func (conn *Connection) Close(){
	if conn.isClose== true{
		return
	}
	conn.isClose=true
	// 线程安全，可重入的close
	conn.Conn.Close()
	close(conn.exitchan)
	close(conn.inchan)
	close(conn.outchan)
}


func (conn *Connection) readLoop(){
	defer conn.Close()
	for{
		_,data,err:=conn.Conn.ReadMessage()
		if err!=nil{
			conn.exitchan<-true
			conn.isClose=true
		}
		conn.inchan <-data
	}
}
func (conn *Connection) writeLoop(){
	for{
		select{
		case data:=<-conn.outchan:
			if err:=conn.Conn.WriteMessage(websocket.TextMessage,data);err!=nil{
				conn.extchan <-true
				return
			}
		case <-conn.exitchan:
			return
		}
	}
}