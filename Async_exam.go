package main

import (
	"fmt"
)

func main() {
	go work1()
	go work2()
	fmt.Println("[全部完成]")
}
func work1() {
	fmt.Println("work1")
}
func work2() {
	fmt.Println("work2")
}
