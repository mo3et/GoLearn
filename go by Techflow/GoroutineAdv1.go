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


// channel传递

func main() {

}
