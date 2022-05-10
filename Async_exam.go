package main

import (
	"fmt"
	"sync"
)

// 不太恰当的示例 体现出不等work1 2就继续执行了
// func main() {
// 	go work1()
// 	go work2()
// 	fmt.Println("[全部完成]")
// 	time.Sleep(100000)
// }
// func work1() {
// 	fmt.Println("work1")
// }
// func work2() {
// 	fmt.Println("work2")
// }

// 正确例子要使用WaitGroup
var waitGroup sync.WaitGroup

func main() {
	waitGroup.Add(2)
	go work1()
	go work2()
	waitGroup.Wait()
	fmt.Println("[全部完成]")
}
func work1() {
	fmt.Println("work1")
	waitGroup.Done()
}
func work2() {
	fmt.Println("work2")
	waitGroup.Done()
}
