package main

import (
	"fmt"
)

// channel Select机制

// 轮询这些Channel 哪个好了就读哪个 否则阻塞等待
// Select 实现
select {	//case 后面必须接一个chan的操作 可以是 input <-chan 也可以是 output <-chan
	case <- chan1:
	case chan2 <-1:
	default: //非必须 
	// 都失败进default 没有default, select会陷入阻塞
	// select是为了从多个数据源中虎丘数据，多个chan同时有数据是，key避免判断哪个数据源数据ready的问题 
}

	// range机制

// range遍历 array or slice
arr:=make([]int, 0)
for i:=range arr{
	// do something
}

map :=make(map[string]int)

for k,v:= range mp{
	// do something
}
// range的底层原理是 chan，当我们使用range时，会不断从chan中接收值，直到关闭

// 可以这样遍历一个chan中的数据
ch:=make(chan int)

for c:=range ch {
	// do something 
}
// End ==================


	// 超时机制
// chan 写入写出都是阻塞的 如往没有缓冲区的chan写入数据需要下游消费才能写入
// 阻塞往往有很大隐患,处理不好容易整个程序锁死
// 需要设计机制来解决 较好方案是设置定时器 Timeout 超时 chan 未响应成功，就人工停止程序

//启动定时器 手动终止Goroutine 结合select机制实现
timeout:=make(chan bool)
go func(){
	time.Sleep(le9)
	timeout<-ture
} ()

select {
	case<-ch:
		// do something
	case <-timeout:
		// break
	}

	//TODO channel传递 
// chan可以传输任何类型的数据 也可以用一个chan 传输chan
// 相当于直接把数据传输给下游 传输读取数据的chan可能更加方便
// 结合函数式编程 把处理的数据一并传输给下游，下游读取数据，并用读取到的处理函数来处理
// 这样更加定制化 以后数据和处理方式都发生变动，也只需在上游更改，可以更加解耦

// chan传递demo
type MetaData struct {
	value interface{}
	handler func(interface{}) int
	downstream chan interface{}
}

func handle (queue chan *MetaData){
	for data:=range queue{
		data.downstream <-data.handler(data.value)
	}
}

	// 单向channel
// 指定channel是只读 or 只写。但数据不可能不完成数据流通 ,并不是真正意义的单向
// 只是为了规范 对使用方进行限制 比如限定在消费函数中不能写入，在生产函数中不能消费
// 在通过函数传递chan时，可以通过加上限定让chan在函数变成单向的

var ch chan <- float32 //只写chan
var ch  <- chan  float32 //只读chan

// 把正常的chan转化成单向chan
var ch chan int
ch 1:= <- chan int (ch) //chan 变成 只读chan
ch 2:=chan <- int (ch)  //chan 变成 只写chan

// 单chan的转化  一般用在函数参数当中 主要是为了起到规范作用
func Test(ch <-chan int) {
	for val:=range ch{
		// do something
	}
} 

func main() {

}
