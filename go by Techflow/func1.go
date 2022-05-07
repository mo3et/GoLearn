package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	defer1()
	defer2()
}

// 变量简写 var a,b int
func add(x, y int) int { //
	return x + y
}

//多值返回
func sample() (string, string) {
	return "sample1", "sample2"
}

//对返回值命名
func sample1(x, y, z int) (xPrime, yPrime, zPrime int) {
	xPrime, yPrime, zPrime = x-1, y+1, z-2
	return
}

// defer 先存入栈 等到函数退出时执行
func defer1() {
	defer fmt.Println("World") //defer 等到其他执行退出了再执行

	fmt.Println("hello")
}

//由于Defer修饰的代码会放入栈中 所以会按照FILO原则
func defer2() {
	for i := 0; i < 10; i++ {
		defer fmt.Print(i) //因为是Stack Type,So output 9876543210
	}
	fmt.Println("done.")
}
