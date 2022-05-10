package main

import (
	"fmt"
	"time"
)

func main() {
	go work1()
	go work2()
	fmt.Println("[全部完成]")
	time.Sleep(1000)
}
func work1() {
	fmt.Println("work1")
}
func work2() {
	fmt.Println("work2")
}
