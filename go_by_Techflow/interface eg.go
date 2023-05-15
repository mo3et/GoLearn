package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I
	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	interf_eg()
	nil_tst()
}

func describe(i I) {
	fmt.Printf("(%v,%T)\n", i, i)
}

// 把interface理解成一种万能数据类型 可以接收任何值
func interf_eg() {
	//定义interface{}时 其实是定义了空的interface
	//其实是定义了空的interface
	var a1 interface{} = 1
	var a2 interface{} = "abc"
	list := make([]interface{}, 0) //构建一个interface{}类型的slice 可以接收任何类型的值和实例
	list = append(list, a1)
	list = append(list, a2)
	fmt.Println(list)

	// 判断 interface的变量类型
	// 在interface变量后面用 .(type) 判断
	if v, ok := a1.(int); ok {
		fmt.Println(v)
	}

	// 类型较多用 switch实现
	switch v := a2.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println(v)
	}
}

// 空值nil

type IN interface {
	MN()
}

type TN struct {
	S string
}

func (t *TN) MN() {
	// 如果不判断nil 会引发一个 空指针进行寻址 err
	if t == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println(t.S)
}

// i为nil invalid memory address or nil pointer dereference
func nil_tst() {
	var i IN
	var t *TN
	i = t
	i.MN()
}

/* //TODO： nil 触发异常是最常见的问题
也要求我们在实现结构体内方法的时候
一定要判断调用的对象是否为nil */


// 赋值类型选择

// e.g
type Integer interface

// Less 和 Add 针对的类型是不同的
// Less不需要修改原值 所以传入 Integer值 
// Add 需要修改原值 所以传入的类型是 Integer的指针
type Operation interface {
	Less(b Integer) bool
	Add(b Integer)
}

func (a Integer) Less(b Integer) bool {
	return a<b
}
func (a *Integer) Add(b Integer) {
	*a+=b
}

func interfacerun(){
	var a Integer =1
	var b Operation = &a //传入指针 会自动生成新的Less方法 做中转
	// var b Operation = a  
}

// Summary：
// Interface 作为万能类型接收各种格式的值。
// interface 空指针调用问题
// Interface 中的两个函数接收类型不一致的问题