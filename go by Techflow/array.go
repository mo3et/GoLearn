package main

import (
	"fmt"
)

func main() {
	//声明一个长度为10的int型的数组
	var a [10]int //定义长度后不能改变
	slice2()
}

func slice1() {
	var a [10]int
	var s []int = a[0:4]

	//也可以直接利用数组给切片赋值
	s := a[:4]
}

//切片 slice 只是对数组的一个引用 如果原数组发生变化 则切片中的数据会一同变化
func slice2() {
	var a [10]int
	var s []int = a[0:4]
	fmt.Println(s) //Output [0 0 0 0 ]
	a[0] = 4
	fmt.Println(s) //Output [4 0 0 0 ]
}
