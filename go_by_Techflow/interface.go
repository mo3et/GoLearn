package main

import (
	"fmt"
)

// Interface的本质都是为了抽象 通过接口提取出一些方法
// Interface就是只有抽象方法(无法直接创建实例)的抽象类

// 创建interface
type Mammal interface {
	Say()
}

// 只要拥有这个函数的结构体 就可以用这个接口来接收

type Dog struct{}

type Cat struct{}

type Human struct{}

func (d Dog) Say() {
	fmt.Println("woof")
}
func (c Cat) Say() {
	fmt.Println("meow")
}
func (h Human) Say() {
	fmt.Println("speak")
}

// 使用接口接收各种结构体对象 然后调用方法

func main() {
	var m Mammal
	m = Dog{}
	m.Say()
	m = Cat{}
	m.Say()
	m = Human{}
	m.Say()
}

// Summary：
// golang的接口 解耦了接口和实现类直接的联系 解决了供需关系颠倒的问题
