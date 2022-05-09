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

func main() {

}
