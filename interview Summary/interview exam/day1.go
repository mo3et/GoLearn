// Articles: https://www.topgoer.cn/docs/gomianshiti/mian1
package main

import (
	"fmt"
)
func main() {
	main1()
	main2()
	main5()
}


//  Day 1



func main1() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

// 参考解析：defer 的执行顺序是后进先出。当出现 panic 语句的时候，会先按照 defer 的后进先出的顺序执行，最后才会执行panic
// ======================

// Day 2
func main2() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		// m[key] = &val  //wrong code
		value := val
		m[key] = &value
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

// Output:
/*
0 -> 3
1 -> 3
2 -> 3
3 -> 3
*/


// Day 3
// 一.
// 1.
func main3_1() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}// Output [0 0 0 0 0 1 2 3]

// 2.
func main3_2() {
   s := make([]int,0)
   s = append(s,1,2,3,4)
   fmt.Println(s)
}// Output [1 2 3 4]

// 二.
func funcMui(x,y int)(sum int,error){
	return x+y,nil
}
 //缺陷  第二个返回值没有命名

//  三.

// new 与 make的区别

// new(T) 和 make(T,args) 是 Go 语言内建函数，用来分配内存，但适用的类型不同。

// new(T) 会为 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T的值。换句话说就是，返回一个指针，该指针指向新分配的、类型为 T 的零值。适用于值类型，如数组、结构体等。

// make(T,args) 返回初始化之后的 T 类型的值，这个值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。make() 只适用于 slice、map 和 channel.

// Day 4

// 1.
func main4_1(){
	list:=new([]int)//不能用new 不然生成是 *[]int 类型的指针 不能对指针执行 append11
	list:=append(list,1)
	fmt.Println(list)
}
// 参考答案及解析：
// 不能通过编译，new([]int) 之后的 list 是一个 *[]int 类型的指针，不能对指针执行 append 操作。可以使用 make() 初始化之后再用。同样的，map 和 channel 建议使用 make() 或字面量的方式初始化，不要用 new() 。

// 2. 
func main4_2() {
	s1:=[]int{1,2,3}
	s2:=[]int{4,5}
	s1:=append(s1,s2)//需要改为s2...
	fmt.Println(s1)
}//不能通过编译 append 第二个参数不能直接使用slice 需要... 不定参数


// 3.
var {
	size:=1024//只能在函数内部使用简短模式
	max_size=size*2
}
func main4_3() {
	fmt.Println(size,max_size)
}
/*
	使用限制
1.必须使用显示初始化
2.不能提供数据类型，编译器会自动推导
3.只能在函数内部使用简短模式
 */

//  Day 5
func main5() {
	sn1:=struct{
		age int
		name string
	}{age:11 ,name: "qq"}
	sn2:=struct{
		age int
		name string
	}{age:11,name: "qq"}

	if sn1==sn2{
		fmt.Println("sn1 == sn2")
	}
	sm1:=struct{
	age int
		m map[string]string
	}{age:11,m:map[string]string{"a":"1"}}
	
	sm2:=struct{
	age int
		m map[string]string
	}{age:11,m:map[string]string{"a":"1"}}
if sm1==sm2{
	fmt.Println("sm1 == sm2")
}	
}