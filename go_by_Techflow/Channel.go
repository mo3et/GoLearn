package main

import (
	"fmt"
)

// Channel
// chan 用来 在goroutine之间传输数据
// 通过make 创建 (slice map channel)

Ch:=make(chan int) //chan 后面为信道 传输的数据类型. 传输任何类型：用interface{}

// 放入数据用 <- 
// 用箭头表示数据的流动

Ch <- a //放入数据到chan  
v:=<-Ch //从chan获取数据

	//	阻塞 
// chan的使用是阻塞的 
// 只有下游从chan中拿走一个数据,我们才能传入一个数据.否则,传输数据的代码会一直等待chan清空
// 同样 如果定义了一个从chan中读取数据的语句，假如chan当前为空，那么会一直阻塞等待，直到chan中有数据

// chan的使用场景中需要一个生产方，也需要一个消费方 


// ===================================
// golang官方例子
package main

 import "fmt"

 func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)
}
// Tips：
// c<- sum 是chan数据写入   x<-c 是chan数据输出 取消阻塞
//chan的传输是阻塞，所以这一语句会一直等待 知道上面两个Goroutine都计算完成为止 
// ===================================

	// chan的缓冲
// chan的容量只有1 只有消费端和生产端都就绪时才可以传输数据
// 我们可以给chan加上缓冲 消费端来不及时 允许生产端先把数据暂存chan，先不发生堵塞 只有chan满了才堵塞

// Chan buffer Usage
Ch:=make(chan int,100)	//创建了缓冲区为100的channel
// Tips:
// 这种情况不常用 因为 上下游的消费情况是统一的 如生产速度过快，消费跟不上，先暂存缓冲区也没太大作用，早晚堵塞

	// close
// 使用channel结束后，可以用close关闭
// Close 只能在生产端进行, 消费端 Close 会引发panic
// 例如在 chan <-input 里 加上Close
// 从接收数据时,可以加上一个参数判断channel是否关闭

// 在消费端 output <-chan 判断是否关闭
v,ok:=<-chan
if !ok {
	return
}

func main() {

}
