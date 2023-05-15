package main

import (
	"fmt"
)

func main() {
	//声明一个长度为10的int型的数组
	var a [10]int //定义长度后不能改变
	axa := []int{}
	fmt.Println(a, axa)
	slice2()

	//标准声明写法
	b := []int{1, 2, 3, 4, 5, 6}
	var xs []int = b[0:4]
	printSlice(xs)

	s := []int{1, 2, 3, 4, 5, 6}
	printSlice(s)

	//对这个切片再切片
	s = s[:2]
	printSlice(s)
	s = s[:4]
	printSlice(s)

	//另一种方法
	s = s[2:]
	printSlice(s)

	SliceTest()

	// makeSlice()

	appendSlice()

	matSlice()
}

func slice1() {
	//标准写法
	var a [10]int
	var s []int = a[0:4]

	//也可以直接利用数组给切片赋值
	ss := a[:4]
	fmt.Println(s, ss)
}

//切片 slice 只是对数组的一个引用 如果原数组发生变化 则切片中的数据会一同变化
func slice2() {
	var a [10]int
	var s []int = a[0:4]
	fmt.Println(s) //Output [0 0 0 0 ]
	a[0] = 4
	fmt.Println(s) //Output [4 0 0 0 ]
}

//Slice 进阶用法
func slicefinal() {
	var a = [3]int{0, 1, 2} //这是个数组声明，固定长度，并用指定元素进行初始化

	//变成切片的声明
	var s = []int{0, 1, 2} //这样在内部创建了数组 我们获取的是数组的切片，数组对我们不可见
	fmt.Println(a)
	fmt.Println(s)
}

/**长度和容量 length 和 capability
长度指的是 切片本身包含的元素的个数  容量是 切片对应的数组从头到末尾包含的元素个数
e.g.
s := []int{1, 2, 3, 4, 5, 6,7,8}	 //cap为8
s= s[:4]	//len为4
**/
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func SliceTest() {
	var a = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println()
	//full
	s := a[:]
	printSlice(s)

	s = s[:2]
	printSlice(s)
	s = s[2:]
	printSlice(s)
	//s[0] = 4
	fmt.Println(a)

}

// make Slice
func makeSlice() {
	s := make([]int, 0, 5) //长度为0 容量为5 的切片
	fmt.Println(s)
	s1 := make([]int, 5) //只传入一个参数 表示切片的长度和容量相等
	fmt.Println(s1)
}

// append Slice 向切片追加元素
func appendSlice() {
	s := make([]int, 4)
	s = append(s, 4)
	fmt.Println("Append Slice:")
	fmt.Println(s)
}

// 二维切片
func matSlice() {
	mat := make([][]int, 10)
	fmt.Println("matSlice:")
	for i := 10; i < 10; i++ {
		mat[i] = make([]int, 10)
		fmt.Println(mat[i])
	}
}
