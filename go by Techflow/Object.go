package main

import (
	"fmt"
	"math"
)

// struct 例子  定义了Point 结构体 和面向结构体的方法Dis
type Point struct {
	x int
	y int
}

// 面向结构体Point的函数 Dis
// 在func 和 函数名之间多了一个(p Point) 是 定义这个函数的接收者
func (p Point) Dis() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type Integer int //从int定义了新的类型 Integer

func (a Integer) Less(b Integer) bool {
	return a < b
}

// Tips:
// 	Golang不允许给简单的内置类型添加方法，并接收者的类型定义和方法声明必须在同一个包里

// 指针接收者
// 我们可以将函数的接收者 定义成指针类型
type Point2 struct {
	x int
	y int
}

// 改写成指针类型
func (p *Point2) Dis2() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

// 如果是指针类型 可以在函数内修改这个结构体的值
func (p *Point2) Modify1() {
	p.x += 5
	p.y -= 3
}
func main2() {
	p3 := Point2{3, 4}
	p3.Modify1()
	fmt.Print(p3)
}

func main() {
	// p1 := Point{3, 4}
	// fmt.Print(p1.Dis())
	// fmt.Print(p1.Less())

	// 指针接收者的使用
	p2 := Point2{3, 4}
	fmt.Print(p2.Dis2())

	main2()
}

//Summary:
//struct 是golang中面向对象的载体

/* 对于结构体来说，我们可以把它当做是接受者传递给一个函数，使得我们可以以类似调用类
当中方法的形式来调用一个函数。并且对于函数而言，接受者除了值以外还可以是一个指
针。如果是指针的话，当我们对结构体值进行修改的时候，会影响到原值。即使我们定义的
接收者类型是指针，我们在调用的时候也不必显示将它转化成结构体指针，golang当中会自
动替我们完成这样的转化。 */
