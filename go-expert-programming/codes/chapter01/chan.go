package main

import (
    "fmt"
)

// # channel data Structure
// `src/runtime.chan.go:hchan` 定义了Channel的数据结构
type hchan struct{
    qcount uint // 当前队列中剩余元素个数
    dataqsiz uint // 环形队列长度,即可以存放的元素个数
    buf unsafe.Pointer //环形队列指针
    elemsize uint16 // 每个元素的大小
    closed uint32 // 表示关闭状态
    elemtype *_type //元素类型
    sendx uint // 队列下标,指示元素写入时存放到队列中的位置
    recvx uint // 队列下标,指示元素从队列的该位置读出
    recvq waitq // 等待读消息的goroutine 队列
    sendq waitq // 等待写消息的goroutine 队列
    lock mutex  // 互斥锁,chan不允许并发读写 
}
// 从数据结构可以看出, Channel 由队列,类型信息,goroutine等待队列组成.



// Create channel   (pseudo code)
func makechan(t *chantype, size int) *hchan{
    var c *hchan
    c=new(hchan)
    c.buf=malloc()
}



//func main(){

